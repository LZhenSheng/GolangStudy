package main

import (
	"GolangStudy/Introduction/grpc/protos"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	// req := protos.HelloRquest{
	// 	Name:    "bobby",
	// 	Age:     18,
	// 	Courses: []string{"go", "gin", "微服务"},
	// }
	// jsonStruct := Hello{Name: "bobby"}
	// jsonRsp, _ := json.Marshal(jsonStruct)
	// fmt.Println(len(string(jsonRsp)))
	// rsp, _ := proto.Marshal(&req)
	// fmt.Println(len(string(rsp)))

	req := protos.HelloRquest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req)
	newReq := protos.HelloRquest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
