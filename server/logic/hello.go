package logic

import (
	pb "rpc/struct"

	"golang.org/x/net/context"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
