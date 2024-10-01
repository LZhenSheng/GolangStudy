package main

import (
	"GolangStudy/Introduction/Rpc/example/handler"
	"fmt"
)

func main() {
	//进一步改造rpc调用的代码
	//1.实例化一个
	client := handler.NewHelloServiceClient("tcp", "127.0.0.1:1234")
	var reply string
	err := client.Hello("bobby", &reply)
	if err != nil {
		// panic("调用失败")
		fmt.Println(err)
	}
	fmt.Println(reply)
}
