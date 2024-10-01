package main

import (
	"GolangStudy/Introduction/grpc/protos"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := protos.HelloRquest{
		Name: "bobby",
	}
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
}
