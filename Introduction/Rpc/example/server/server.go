package main

import (
	"net"
	"net/rpc"

	"GolangStudy/Introduction/Rpc/example/handler"
)

func main() {
	//进一步改造rpc调用的代码
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2.注册处理逻辑handler
	_ = handler.RegisterHelloService()
	//3.启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候
	rpc.ServeConn(conn)

}
