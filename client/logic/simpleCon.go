package logic

import (
	"log"
	pb "rpc/struct"

	"context"

	"google.golang.org/grpc"
)

var HelloClient pb.GreeterClient

func InitHello(conn *grpc.ClientConn) {
	if conn == nil {
		log.Panic("HelloCLient is failed,conne is err")
	}
	HelloClient = pb.NewGreeterClient(conn)
}

func SayHelllo(name string) {
	res, err := HelloClient.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Println("hello client sayhello is err:", err)
		return
	}
	log.Println("res info is :", res.GetMessage())
}
