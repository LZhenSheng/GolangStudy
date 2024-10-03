package main

import (
	"GolangStudy/Introduction/grpc/validate/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type customCredential struct{}

func main() {
	var opts []grpc.DialOption

	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//rsp, _ := c.Search(context.Background(), &empty.Empty{})
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:     9999,
		Email:  "bobby@qq.com",
		Mobile: "19999999999",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}
