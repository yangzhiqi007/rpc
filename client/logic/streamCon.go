package logic

import (
	"context"
	"fmt"
	"io"
	pb "rpc/struct"
	"sync"
	"time"

	"github.com/silenceper/pool"
)

/*
   初始化
   min // 最小链接数
   max // 最大链接数
   factory func() (interface{}, error) //创建链接的方法
   close func(v interface{}) error //关闭链接的方法
*/
func InitThread(min, max int, factory func() (interface{}, error), close func(v interface{}) error) (pool.Pool, error) {

	poolConfig := &pool.PoolConfig{
		InitialCap: min,
		MaxCap:     max,
		Factory:    factory,
		Close:      close,
		//链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("Init err=", err)
		return nil, err
	}
	return p, nil
}

// 测试连接
func StreamTest(p pool.Pool) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//获取连接
			v, _ := p.Get()
			client := v.(pb.DataClient)
			info := &pb.UserInfoRequest{
				Uid: 10012,
			}
			GetUserInfo(client, info)
			//归还链接
			p.Put(v)
		}()
		wg.Wait()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//获取连接
			v, _ := p.Get()
			client := v.(pb.DataClient)
			ChangeUserInfo(client)
			//归还链接
			p.Put(v)
		}()
		wg.Wait()
	}
	//获取链接池大小
	current := p.Len()
	fmt.Println("len=", current)
}

//简单模式
func GetUserInfo(client pb.DataClient, info *pb.UserInfoRequest) {
	req, err := client.GetUserInfo(context.Background(), info)
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	fmt.Println("userinfo is ", req.GetAge(), req.GetCount(), req.GetName(), req.GetSex())
}

//双向流模式
func ChangeUserInfo(client pb.DataClient) {
	notes := []*pb.UserInfoResponse{
		{Name: "jim", Age: 18, Sex: 2, Count: 100},
		{Name: "Tom", Age: 20, Sex: 1, Count: 666},
	}
	stream, err := client.ChangeUserInfo(context.Background())
	if err != nil {
		fmt.Println("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				fmt.Println("read done ")
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println("Failed to receive a note : %v", err)
			}
			fmt.Println("get message userinfo:", in)
		}
	}()
	fmt.Println("notes", notes)
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			fmt.Println("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
