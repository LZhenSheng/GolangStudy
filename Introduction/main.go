package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func example1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func example2() {
	//
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]

	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), len(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), len(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), len(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), len(s2))

	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), len(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), len(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), len(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), len(s2))
}
func example3() {
	arr := []int{9, 8, 7, 6}
	for index, value := range arr {
		fmt.Printf("%d => %d\n", index, value)
	}
}
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ready body failed:%v", err)
		return
	}
	//类型转换，将[]byte转换为string
	fmt.Fprintf(w, "read the data:%s \n", string(body))

	//尝试再次读取，发现啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		//不会倒这里
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data ", err)
}
func example4() {
	http.HandleFunc("/read/only", readBodyOnce)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Name string
	Age  int
}

// 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}
func example5() {
	//u是结构体，所以方法调用的时候它数值是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)
	//因为up指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}
	//因为ChangeName的接收器是结构体
	//所以Up的数据还是不会变
	up.ChangeName("Jerry Changed")
	up.ChangeAge(120)
	fmt.Printf("%v \n", up)
}

type Server interface {
	//method post get put delete
	Route(pattern string, handleFunc func(ctx *Context))
	Start(address string) error
}

// sdkHttpServer基于http库实现
type sdkHttpServer struct {
	Name string
}

// Route 注册路由
func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, reques *http.Request) {
		ctx := &Context{
			R: reques,
			W: writer,
		}
		handleFunc(ctx)
	})
}
func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败: %v", err)
	}
}

func main() {
	server := NewHttpServer("test-server")
	// server.Route("/", handler)
	server.Route("/user/signup", SignUp)
	server.Start(":8080")
}
