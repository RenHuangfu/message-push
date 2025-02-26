package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"message-push/client/v1"
)

func main() {
	// 使用 grpc.NewClient 创建连接
	conn, err := grpc.NewClient("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	cli := v1.NewDemoClient(conn)
	res, err := cli.StreamMessage(context.Background())
	if err != nil {
		log.Fatalf("Failed to receive a note : %v", err)
	}
	for {
		in, err := res.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}
		log.Printf("Got message (%s)", in.Message)
	}

}
