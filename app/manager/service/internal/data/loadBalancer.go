package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RenHuangfu/tools/kafka"
	"github.com/pkg/errors"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"math/rand"
	"message-push/app/manager/common/model/entity"
	"message-push/app/manager/common/model/po/ent"
	"message-push/app/manager/common/repo"
	"message-push/app/manager/service/internal/conf"
	"strconv"
	"sync"
	"time"
)

type LoadBalancer struct {
	zk              *zk.Conn
	producers       []kafka.Producer
	topics          []string
	sendersPath     string
	sendersNamePath string
	pushers         []*entity.Pusher
	mu              sync.RWMutex
}

func NewLoadBalancer(c *conf.Bootstrap) repo.LoadBalancer {
	conn, _, err := zk.Connect([]string{c.Data.Zookeeper.Addr}, c.Data.Zookeeper.Timeout.AsDuration())
	if err != nil {
		panic(err)
	}

	producers := make([]kafka.Producer, 0)
	for _, topic := range c.Data.Kafka.Topic {
		producerFirst := kafka.NewKafkaProducer(
			kafka.WithBrokers(c.Data.Kafka.Brokers),
			kafka.WithTopic(topic),
			kafka.WithAck(0),
			kafka.WithAsync())
		if producerFirst == nil {
			panic("nil producer")
		}
		producers = append(producers, producerFirst)
	}

	lb := &LoadBalancer{
		zk:              conn,
		producers:       producers,
		topics:          c.Data.Kafka.Topic,
		sendersPath:     c.Data.Zookeeper.SendersRoot,
		sendersNamePath: c.Data.Zookeeper.SendersName,
		pushers:         make([]*entity.Pusher, 0),
	}
	go lb.watchNodes()
	return lb
}

func (lb *LoadBalancer) SelectProducer(ctx context.Context, p *ent.Message) (kafka.Producer, error) {
	index := p.AppID % len(lb.topics)
	if index < 0 {
		index = -index
	}
	return lb.producers[index], nil
}

func (lb *LoadBalancer) SearchPusher(ctx context.Context, p *entity.SearchPusherParam) (string, error) {
	appId, err := strconv.Atoi(p.AppId)
	if err != nil {
		return "", err
	}

	index := appId % len(lb.topics)
	if index < 0 {
		index = -index
	}

	hostPort, err := lb.selectNode(lb.topics[index])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ws://%s/ws", hostPort), nil
}

func (lb *LoadBalancer) selectNode(topic string) (string, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	// 过滤不可用节点
	validNodes := make([]*entity.Pusher, 0)
	for _, node := range lb.pushers {
		if time.Since(node.LastUpdate) > 30*time.Minute {
			continue // 剔除超时节点
		}
		if node.CPULoad > 0.8 {
			continue // 排除高负载节点
		}
		if node.Topic != topic {
			continue
		}
		validNodes = append(validNodes, node)
	}

	if len(validNodes) == 0 {
		return "", errors.New("no available nodes")
	}

	// 计算动态权重
	totalScore := 0
	scores := make([]int, len(validNodes))
	for i, node := range validNodes {
		// 权重计算公式（可调整系数）
		score := int(float64(node.Weight)*0.6 +
			(1-float64(node.Connections)/1000.0)*30 +
			(1-node.CPULoad)*10)
		if score < 1 {
			score = 1 // 最低保障权重
		}
		scores[i] = score
		totalScore += score
	}

	// 生成随机选择
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(totalScore)
	current := 0
	for i, score := range scores {
		current += score
		if current >= pick {
			return fmt.Sprintf("%s:%d",
				validNodes[i].Host, validNodes[i].Port), nil
		}
	}

	return fmt.Sprintf("%s:%d",
		validNodes[0].Host, validNodes[0].Port), nil // fallback
}

func (lb *LoadBalancer) watchNodes() {
	for {
		children, _, eventCh, err := lb.zk.ChildrenW(lb.sendersNamePath)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// 获取最新节点列表
		lb.refreshNodes(children)

		// 阻塞等待变更事件
		select {
		case event := <-eventCh:
			log.Printf("Zookeeper event: %+v", event)
			// 事件触发后重新刷新节点
		case <-time.After(30 * time.Second):
			// 防止长时间无事件导致状态过期
		}
	}
}

// 刷新节点数据
func (lb *LoadBalancer) refreshNodes(children []string) {
	newNodes := make([]*entity.Pusher, 0, len(children))

	for _, child := range children {
		path := fmt.Sprintf("%s/%s", lb.sendersPath, child)
		data, _, err := lb.zk.Get(path)
		if err != nil {
			continue
		}

		var nodeData entity.PusherJson
		if err := json.Unmarshal(data, &nodeData); err != nil {
			continue
		}

		newNode := &entity.Pusher{
			Host:        nodeData.Host,
			Port:        nodeData.Port,
			Weight:      nodeData.Weight,
			Connections: nodeData.Connections,
			CPULoad:     nodeData.CPULoad,
			Topic:       nodeData.Topic,
			LastUpdate:  time.Unix(nodeData.UpdatedAt, 0),
		}
		newNodes = append(newNodes, newNode)
	}

	// 原子更新节点列表
	lb.mu.Lock()
	lb.pushers = newNodes
	lb.mu.Unlock()
}
