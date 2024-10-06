package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	//制定加载好相对路径
// 	//加载二级目录
// 	router.LoadHTMLGlob("templates/**/*")
// 	// router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
// 	router.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl", gin.H{
// 			"title": "mooc",
// 		})
// 	})
// 	router.GET("/goods/list", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "goods/list.html", gin.H{
// 			"name": "微服务开发",
// 		})
// 	})
// 	router.GET("/users/list", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "users/list.html", gin.H{
// 			"name": "微服务开发",
// 		})
// 	})
// 	router.Run()
// }
