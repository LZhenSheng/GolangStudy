package main

import (
	"net"
	"net/rpc"
)

func add(a, b int) int {
	return a + b
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	company Company
}
type PrintResult struct {
	Info string
	Err  error
}

// func RpcPrintln(employee Employee) PrintResult {
/*
	客户端
	1.建立连接tcp/http
	2.将employee对象序列化成json字符串-序列化
	3.发送json字符串-调用成功后实际上你接受到的是一个二进制的数据
	4.等待服务器发送结果
	5将服务返回的数据解析成PrintResult对象-反序列化
	服务端
	1.监听网络接口80
	2.读取数据-二进制的json数据
	3.对数据进行反序列化Employee对象
	4.开始处理业务逻辑
	5.将处理的结果PrintReuslt发序列化成json二进制数据-序列化
	6.将数据返回
	序列化和反序列化是可以选择的，不一定要采用json、xml、protobuf、msgpack
*/
// }
type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func main() {
	// fmt.Println(add(1, 3))
	//1.这个函数的调用参数如何传递-json（json是一种数据格式的协议）/xml/protobuf/msgpack-编码协议，json不是一个高性能的编码协议
	//现在网络调用有两个端-客户端、应该干嘛？将数据传输到gin
	//gin-服务端，服务端负责解析数据
	// fmt.Println(add(1, 2))
	//将这个打印的工作放在另一个台服务器上，我需要将本地的内存对象struct，这样不行，可行的方式就是将struct序列化成json
	// fmt.Println(Employee{
	// 	Name: "bobby",
	// 	company: Company{
	// 		Name: "imooc",
	// 		Address: "beijing",
	// 	},
	// })
	//远程服务器需要将二进制对象反解成struct对象
	//搞这么麻烦，直接全部使用json去格式化不香嘛？这种做法在浏览器和gin服务之间是可以行，但是如果是一个大型的分布式系统不行

	//get方法http://127.0.0.1:8000/add?a=1&b=2或http://127.0.0.1:8000?method=add&a=1&b=2
	//返回的格式化：json{"data":3}
	//1、callId的问题：r.URL.Path，2、数据的传输协议：url的参数传输协议，3、网络传输协议：http
	// http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
	// 	err := r.ParseForm() //解析参数
	// 	if err != nil {
	// 		panic("error")
	// 	}
	// 	fmt.Println("path:", r.URL.Path)
	// 	a, err := strconv.Atoi(r.Form["a"][0])
	// 	if err != nil {
	// 		panic("transform error")
	// 	}
	// 	b, err := strconv.Atoi(r.Form["b"][0])
	// 	if err != nil {
	// 		panic("transform error")
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	jData, err := json.Marshal(map[string]int{
	// 		"data": a + b,
	// 	})
	// 	w.Write(jData)
	// })
	// _ = http.ListenAndServe(":8000", http.TimeoutHandler(http.DefaultServeMux, time.Duration(10*time.Second), "Request timed out"))

	//rpc快速开发体验
	// //1.实例化一个server
	// listener, _ := net.Listen("tcp", ":1234")
	// //2.注册处理逻辑
	// _ = rpc.RegisterName("HelloService", &HelloService{})
	// //3.启动服务
	// conn, _ := listener.Accept() //当一个新的连接进来的时候
	// rpc.ServeConn(conn)

	//一连串的代码大部分都是net的包好像和rpc没有关系
	//不行。rpc调用中有几个问题需要解决1.call id2.序列化和反序列化 编码和解码
	//可以跨语言调用呢1.go语言的rpc的序列化和反序列化协议是什么（Gob）.2.能否替换成常见的序列化

	// //替换rpc的序列化协议为json
	// //1.实例化一个server
	// listener, _ := net.Listen("tcp", ":1234")
	// //2.注册处理逻辑
	// _ = rpc.RegisterName("HelloService", &HelloService{})
	// //3.启动服务
	// conn, _ := listener.Accept() //当一个新的连接进来的时候
	// rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	// //替换rpc的序列化协议为http
	// //2.注册处理逻辑
	// _ = rpc.RegisterName("HelloService", &HelloService{})
	// http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
	// 	var conn io.ReadWriteCloser = struct {
	// 		io.Writer
	// 		io.ReadCloser
	// 	}{
	// 		ReadCloser: r.Body,
	// 		Writer:     w,
	// 	}
	// 	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	// })
	// http.ListenAndServe(":1234", nil)

	//进一步改造rpc调用的代码
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候
	rpc.ServeConn(conn)

}
