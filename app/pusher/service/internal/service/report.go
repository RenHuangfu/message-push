package service

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/shirou/gopsutil/v3/cpu"
	"message-push/app/pusher/common/model/entity"
	"message-push/app/pusher/service/internal/conf"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Report struct {
	zk *zk.Conn
}

func NewReport(c *conf.Bootstrap) *Report {
	conn, _, err := zk.Connect([]string{c.Data.Zookeeper.Addr}, c.Data.Zookeeper.Timeout.AsDuration())
	if err != nil {
		panic(err)
	}
	r := &Report{
		zk: conn,
	}
	go r.StartStatusReporter(c)
	return r
}

func (r *Report) StartStatusReporter(c *conf.Bootstrap) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			data := &entity.PusherJson{
				Host:        c.Report.Host,
				Port:        int(c.Report.Port),
				Topic:       c.Report.Topic,
				Weight:      int(c.Report.Weight),
				Connections: getConnectionCount(),
				CPULoad:     getCPULoad(),
				UpdatedAt:   time.Now().Unix(),
			}
			jsonData, _ := json.Marshal(data)

			// 原子更新Zookeeper节点数据
			_, stat, _ := r.zk.Get(c.Report.NodePath)
			_, _ = r.zk.Set(c.Report.NodePath, jsonData, stat.Version)
		}
	}
}

// 获取当前TCP连接数（Linux系统示例）
func getConnectionCount() int {
	out, _ := exec.Command("sh", "-c",
		"netstat -an | grep :8080 | grep ESTABLISHED | wc -l").Output()
	count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
	return count
}

// 获取最近1秒内的CPU使用率（所有逻辑核心的平均值）
func getCPULoad() float64 {
	// 获取1秒内的使用率
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil || len(percentages) == 0 {
		return 0.0
	}
	return percentages[0] / 100.0 // 转换为0-1范围
}
