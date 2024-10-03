package main

import (
	"GolangStudy/Introduction/grpc/error/proto"
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRquest) (*proto.HelloReply, error) {
	time.Sleep(5 * time.Second)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}
func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start ")
	}
}
