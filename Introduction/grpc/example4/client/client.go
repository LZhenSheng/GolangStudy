package main

import (
	_ "GolangStudy/Introduction/grpc/example4/proto"
	"GolangStudy/Introduction/grpc/example4/proto1"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// _ = proto.HelloReply_Result{}
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto1.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto1.HelloRquest{
		Name: "bobby",
		G:    proto1.Gender_FEMALE,
		Mp: map[string]string{
			"name":    "bobby",
			"company": "mooc",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
