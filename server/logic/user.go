package logic

import (
	"fmt"
	"io"
	pb "rpc/struct"

	"golang.org/x/net/context"
)

type UserServer struct {
}

//简单模式
func (this *UserServer) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	uid := in.GetUid()
	fmt.Println("The uid is ", uid)
	return &pb.UserInfoResponse{
		Name:  "Jim",
		Age:   18,
		Sex:   0,
		Count: 1000,
	}, nil
}

//双向流模式
func (this *UserServer) ChangeUserInfo(stream pb.Data_ChangeUserInfoServer) error {
	notes := []*pb.UserInfoResponse{}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			break
		}
		if err != nil {
			fmt.Println("stream is err:", err)
			return err
		}
		fmt.Println("userinfo: ", in)
		notes = append(notes, in)
	}
	for _, note := range notes {
		note.Age += 10
		if err := stream.Send(note); err != nil {
			return err
		}
	}
	return nil
}
