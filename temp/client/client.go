package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/kirinlabs/HttpRequest"
// )

// type ResponseData struct {
// 	Data int `json:"data"`
// }

// func Add(a, b int) int {
// 	//http协议
// 	req := HttpRequest.NewRequest()
// 	req.SetTimeout(10 * time.Second)
// 	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
// 	body, _ := res.Body()
// 	rspData := ResponseData{}
// 	_ = json.Unmarshal(body, &rspData)
// 	return rspData.Data
// }
// func main() {
// 	// fmt.Println(Add(2, 2))

// 	// //rpc快速开发体验
// 	// //1.建立连接
// 	// client, err := rpc.Dial("tcp", "localhost:1234")
// 	// if err != nil {
// 	// 	panic("连接失败")
// 	// }
// 	// var reply *string = new(string)
// 	// err = client.Call("HelloService.Hello", "bobby", reply)
// 	// if err != nil {
// 	// 	panic("调用失败")
// 	// }
// 	// fmt.Println(*reply)

// 	// //替换rpc的序列化协议为json
// 	// //1.建立连接
// 	// client, err := rpc.Dial("tcp", "localhost:1234")
// 	// if err != nil {
// 	// 	panic("连接失败")
// 	// }
// 	// var reply *string = new(string)
// 	// err = client.Call("HelloService", "bobby", reply)
// 	// if err != nil {
// 	// 	panic("调用失败")
// 	// }
// 	// fmt.Println(*reply)

// 	//进一步改造rpc调用的代码
// 	//1.实例化一个
// 	client := NewHelloServiceClient("tcp", "127.0.0.1:1234")
// 	var reply string
// 	err := client.Hello("HelloService.Hello", "bobby", &reply)
// 	if err != nil {
// 		// panic("调用失败")
// 		fmt.Println(err)
// 	}
// 	fmt.Println(reply)
// }
