package main

import (
	"GolangStudy/Introduction/grpc/stream_grpc_test/proto"
	"fmt"
	"net"
	"sync"
	"time"

	_ "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.SteramResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PostStream(cliStr proto.Greeter_PostStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.SteramResData{
				Data: "我是服务器",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}
func main() {
	// empty.Empty{}
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic("failed to start ")
	}
}
