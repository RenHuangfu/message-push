package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/tx7do/kratos-transport/broker"
	"message-push/app/pusher/common/model/entity"
	"strconv"
	"sync"
	"time"
)

type SendService struct {
	log        *log.Helper
	clientConn sync.Map
	report     *Report
}

func NewSendService(report *Report, logger log.Logger) *SendService {
	return &SendService{
		log:    log.NewHelper(logger),
		report: report,
	}
}

func (s *SendService) Consume(ctx context.Context, _ string, headers broker.Headers, msg *entity.Message) error {
	s.log.Infof("Consume: %v", msg)
	clientKeys := make([]entity.ClientKey, len(msg.ClientIDs))
	for i, clientID := range msg.ClientIDs {
		clientKeys[i] = entity.ClientKey{
			AppId:    strconv.Itoa(msg.AppID),
			ClientId: strconv.Itoa(clientID),
		}
	}
	for _, clientKey := range clientKeys {
		go func(clientKey entity.ClientKey) {
			t := time.NewTimer(time.Hour * 24)
			for {
				select {
				case <-t.C:
					return
				default:
					if newConn, ok := s.clientConn.Load(clientKey); ok {
						conn := newConn.(*entity.Connect)
						conn.Lock()
						err := conn.Conn.WriteMessage(websocket.TextMessage, []byte(msg.Content))
						conn.Unlock()
						if err != nil {
							s.log.Infof("send message failed: %v", err)
							s.clientConn.CompareAndDelete(clientKey, newConn)
						} else {
							return
						}
					}
					time.Sleep(time.Second * 5)
				}
			}
		}(clientKey)
	}
	return nil
}

func (s *SendService) RegisterClient(req *entity.ClientRequest, conn *websocket.Conn) {
	clientKey := entity.ClientKey{
		AppId:    req.AppId,
		ClientId: req.ClientId,
	}

	if existing, ok := s.clientConn.Load(clientKey); ok {
		s.log.Infof("关闭旧连接: %v", clientKey)
		existingConn := existing.(*entity.Connect)
		existingConn.Lock()
		existingConn.Conn.Close()
		existingConn.Unlock()
	}

	newConn := &entity.Connect{
		Conn: conn,
	}
	s.clientConn.Store(clientKey, newConn)

	ctx := context.WithValue(context.Background(), "clientKey", clientKey)
	cancelCtx, cancel := context.WithCancel(ctx)

	go s.WriteLoop(ctx, cancel, newConn)
	go s.ReadLoop(cancelCtx, newConn)
}

func (s *SendService) ReadLoop(ctx context.Context, conn *entity.Connect) {
	defer func() {
		conn.Lock()
		conn.Conn.Close()
		conn.Unlock()
	}()
	for {
		select {
		case <-ctx.Done():
			s.log.Infof("read loop exit")
			return
		default:
			msgType, data, err := conn.Conn.ReadMessage()
			if err != nil {
				s.log.Errorf("read message failed: %v", err)
				return
			}
			s.log.Infof("received message: type=%d, data=%s", msgType, data)
		}
	}
}

func (s *SendService) WriteLoop(ctx context.Context, cancel context.CancelFunc, conn *entity.Connect) {
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
		clientKey := ctx.Value("clientKey").(entity.ClientKey)
		s.clientConn.CompareAndDelete(clientKey, conn)
	}()
	for {
		select {
		case <-ticker.C:
			conn.Lock()
			err := conn.Conn.WriteControl(
				websocket.PingMessage,
				[]byte("ping"),
				time.Now().Add(5*time.Second),
			)
			conn.Unlock()
			if err != nil {
				s.log.Errorf("心跳发送失败: %v", err)
				cancel()
				return
			}
		}
	}
}
