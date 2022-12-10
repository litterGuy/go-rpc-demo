package main

import (
	"context"
	"fmt"
	pb "go-rpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	// grpc.WithInsecure() 跳过了对服务器证书的验证
	// 连接gRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to dial to: " + Address)
	}
	if conn == nil {
		panic("Failed to get connection from the server")
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Printf("Fail to call method, err: %v", err)
	}

	fmt.Println(r.Message)
}
