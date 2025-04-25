package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/tx7do/kratos-transport/broker"
	"message-push/app/pusher/common/model/entity"
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

func (s *SendService) Consume(ctx context.Context, _ string, _ broker.Headers, _ *entity.Message) error {
	return nil
}

func (s *SendService) RegisterClient(req *entity.ClientRequest, conn *websocket.Conn) {
	s.clientConn.Store(entity.ClientKey{
		AppId:    req.AppId,
		ClientId: req.ClientId,
		Type:     req.Type,
	}, conn)

	ctx := context.WithValue(context.Background(), "clientKey", entity.ClientKey{
		AppId:    req.AppId,
		ClientId: req.ClientId,
		Type:     req.Type,
	})
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
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 发送心跳消息
			if err := conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
				s.log.Errorf("send heartbeat failed: %v", err)
				s.clientConn.Delete(ctx.Value("clientKey"))
				cancel()
				return
			}
		}
	}
}
