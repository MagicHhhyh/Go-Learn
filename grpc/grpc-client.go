package main

import (
	service "GoLearn/grpc/auto-grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 代码段：创建gRPC客户端并连接到服务器，执行一个RPC调用

	// 1. 连接到gRPC服务器，这里使用了不安全的连接方式（禁用加密）
	// 使用grpc.WithTransportCredentials(insecure.NewCredentials())来指定不使用TLS加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// 如果连接服务器失败，则记录错误日志
		log.Printf("无法连接到服务器: %v\n", err)
		return // 发生错误时退出函数
	}
	// 使用defer语句确保在函数退出前关闭连接
	defer conn.Close()

	// 2. 创建客户端对象，这里假设service包中有一个NewHelloServiceClient函数
	// 该函数接受一个连接对象并返回一个客户端实例
	client := service.NewHelloServiceClient(conn)

	// 3. 执行RPC调用，这里调用的是客户端的SayHello方法
	// SayHello方法需要一个上下文对象和一个HelloRequest消息作为参数
	// 这里假设HelloRequest消息中有一个Name字段，我们将其设置为"焦糖玛奇朵"
	resp, err := client.SayHello(context.Background(), &service.HelloRequest{Name: "焦糖玛奇朵"})
	if err != nil {
		// 如果执行RPC调用时发生错误，则记录错误日志
		log.Printf("执行SayHello调用出错: %v\n", err)
		return // 发生错误时退出函数
	}

	// 4. 输出响应结果，这里假设HelloResponse消息中有一个GetMessage方法
	// 该方法返回一个字符串，我们将其打印到控制台
	fmt.Println(resp.GetMessage())
}
