package main

import (
	"GolangStudy/Introduction/grpc/stream_grpc_test/proto"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//服务端流模式
	// c := proto.NewGreeterClient(conn)
	// res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "mooc"})
	// for {
	// 	a, err := res.Recv() //socket编程send recv
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Println(a)

	// }

	// //客户端流模式
	// c := proto.NewGreeterClient(conn)
	// putS, _ := c.PostStream(context.Background())
	// i := 0
	// for {
	// 	i++
	// 	_ = putS.Send(&proto.StreamReqData{
	// 		Data: fmt.Sprintf("mooc%d", i),
	// 	})
	// 	time.Sleep(time.Second)
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	//双向流模式
	c := proto.NewGreeterClient(conn)
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务器消息:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamReqData{
				Data: "我是客户端",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
