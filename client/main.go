package main

import (
	"log"
	"rpc/client/logic"

	pb "rpc/struct"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	// 简单模式
	logic.InitHello(conn)
	logic.SayHelllo("world")

	// 双向流模式
	// 创建连接
	factory := func() (interface{}, error) {
		return pb.NewDataClient(conn), nil
	}
	// 关闭链接，此处只是定义不需要调用了因为上面有defer conn.Close()，定义的目的在于初始化链接池。
	close := func(v interface{}) error { return conn.Close() }
	//初始化链接池
	p, err := logic.InitThread(10, 30, factory, close)
	if err != nil {
		log.Panic("init stream Thread error")
		return
	}
	logic.StreamTest(p)
	select {}
}
