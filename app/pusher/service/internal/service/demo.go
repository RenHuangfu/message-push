package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	v1 "message-push/app/pusher/service/api/server/v1"
	"sync"
)

type Demo struct {
	v1.UnimplementedDemoServer
	log              *log.Helper
	localConnections sync.Map
}

func NewDemo(logger log.Logger) *Demo {
	return &Demo{
		log: log.NewHelper(logger),
	}
}

func (d *Demo) SendMessage(ctx context.Context, request *v1.SendMessageRequest) (*v1.SendMessageResponse, error) {
	if stream, ok := d.localConnections.Load(1); ok {
		// 类型断言为gRPC流
		pushStream, _ := stream.(v1.Demo_StreamMessageServer)
		// 发送消息
		if err := pushStream.Send(&v1.SendMessageResponse{
			Message: request.Message,
		}); err != nil {
			return nil, err
		}
	}
	return &v1.SendMessageResponse{}, nil
}

func (d *Demo) StreamMessage(conn v1.Demo_StreamMessageServer) error {
	d.localConnections.Store(1, conn)
	for {
		_, err := conn.Recv()
		if err == io.EOF {
			// 流结束
			fmt.Println("客户端发送的数据流结束")
			return nil
		}
		if err != nil {
			// 流出现错误
			fmt.Println("接收数据出错:", err)
			return err
		}
	}
}
