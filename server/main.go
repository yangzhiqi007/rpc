package main

import (
	"log"
	"net"
	"rpc/server/logic"
	pb "rpc/struct"

	"google.golang.org/grpc"
)

const Port = ":50051"

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &logic.HelloServer{})
	pb.RegisterDataServer(s, &logic.UserServer{})
	s.Serve(lis)
}
