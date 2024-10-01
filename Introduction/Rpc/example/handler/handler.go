package handler

import "net/rpc"

const HelloServiceName = "HelloServiceName"

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protcol, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{conn}
}
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}

type NewHelloService struct {
}
type HelloServicer interface {
	Hello(request string, reply *string) error
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func RegisterHelloService() error {
	return rpc.RegisterName(HelloServiceName, &NewHelloService{})
}
