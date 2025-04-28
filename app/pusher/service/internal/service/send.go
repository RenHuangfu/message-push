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
}

func NewSendService(logger log.Logger) *SendService {
	return &SendService{
		log: log.NewHelper(logger),
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
					if conn, ok := s.clientConn.Load(clientKey); ok {
						if err := conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(msg.Content)); err != nil {
							s.log.Errorf("send message failed: %v", err)
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
		existing.(*websocket.Conn).Close()
		s.clientConn.Delete(clientKey)
	}

	s.clientConn.Store(clientKey, conn)

	ctx := context.WithValue(context.Background(), "clientKey", clientKey)
	cancelCtx, cancel := context.WithCancel(ctx)

	go s.WriteLoop(ctx, cancel, conn)
	go s.ReadLoop(cancelCtx, conn)
}

func (s *SendService) ReadLoop(ctx context.Context, conn *websocket.Conn) {
	defer conn.Close()
	for {
		select {
		case <-ctx.Done():
			s.log.Infof("read loop exit")
			return
		default:
			msgType, data, err := conn.ReadMessage()
			if err != nil {
				s.log.Errorf("read message failed: %v", err)
				return
			}
			s.log.Infof("received message: type=%d, data=%s", msgType, data)

			if err := conn.WriteMessage(msgType, data); err != nil {
				s.log.Errorf("write message failed: %v", err)
				return
			}
		}
	}
}

func (s *SendService) WriteLoop(ctx context.Context, cancel context.CancelFunc, conn *websocket.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
		clientKey := ctx.Value("clientKey").(entity.ClientKey)
		s.clientConn.CompareAndDelete(clientKey, conn)
	}()
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteControl(
				websocket.PingMessage,
				[]byte("ping"),
				time.Now().Add(5*time.Second),
			); err != nil {
				s.log.Errorf("心跳发送失败: %v", err)
				cancel()
				return
			}
		}
	}
}
