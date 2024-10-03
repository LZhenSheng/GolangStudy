package main

import (
	"GolangStudy/Introduction/grpc/example2/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRquest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "helo" + request.Name,
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
