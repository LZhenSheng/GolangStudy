package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	//优雅推出，当我们关闭程序后需要做的处理
	//微服务 启动之前或者启动之后必须做一件事：将当前的ip地址和端口号注册到注册中心
	//我们当前的服务停止之后并没有告知注册中心
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	go func() {
		router.Run()
	}()
	//接受信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//处理后续的逻辑
	fmt.Println("关闭server中。。。")
	fmt.Println("注册服务。。。")
}
