package main

import (
	"GoLearn/grpc/test1/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func (server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {

	return &proto.HelloResponse{Message: string("我想发送的信息：" + req.Name + "你好！")}, nil
}

func main() {
	// 代码段：启动gRPC服务器并监听指定端口

	// 1. 开启TCP端口监听，指定监听本地的9090端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		// 如果监听端口时发生错误，则记录错误日志
		log.Printf("无法开启端口监听: %v\n", err)
		return // 发生错误时退出函数
	}

	// 2. 创建一个新的gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 3. 注册gRPC服务到服务器实例中
	// 假设我们有一个名为service的包，其中包含RegisterHelloServiceServer函数
	// 该函数将HelloServiceServer的实现注册到gRPC服务器中
	proto.RegisterHelloServiceServer(grpcServer, &server{})

	// 4. 启动gRPC服务器，开始监听并处理来自客户端的请求
	// 服务器将使用之前创建的监听器来接收连接
	if err := grpcServer.Serve(listen); err != nil {
		// 如果启动服务器时发生错误，则记录错误日志
		log.Printf("无法启动gRPC服务器: %v\n", err)
		return // 发生错误时退出函数
	}

	// 服务器启动后，将一直运行，直到程序被外部中断或调用Stop方法
}
