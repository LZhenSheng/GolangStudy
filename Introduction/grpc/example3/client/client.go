package main

import (
	"GolangStudy/Introduction/grpc/example2/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRquest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
