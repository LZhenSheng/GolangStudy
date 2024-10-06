package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func welcome(c *gin.Context) {
// 	firstName := c.DefaultQuery("firstname", "bobby")
// 	lastName := c.Query("lastname")
// 	c.JSON(http.StatusOK, gin.H{
// 		"first_name": firstName,
// 		"last_name":  lastName,
// 	})
// }

// func formPost(c *gin.Context) {
// 	message := c.PostForm("message")
// 	nick := c.DefaultPostForm("nick", "anonymous")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": message,
// 		"nick":    nick,
// 	})
// }

// func main() {
// 	router := gin.Default()
// 	//get请求
// 	router.GET("/welcome", welcome)
// 	//post请求
// 	router.POST("/form", formPost)
// 	//post & get
// 	router.POST("/postsd", func(c *gin.Context) {
// 		id := c.Query("id")
// 		page := c.DefaultQuery("page", "0")
// 		name := c.PostForm("name")
// 		message := c.PostForm("message")

// 		fmt.Printf("id: %s,page: %s,name: %s,message: %s", id, page, name, message)
// 	})
// 	router.Run() // listen and serve on 0.0.0.0:8080
// 	// r.Run(":8083") // listen and serve on 0.0.0.0:8080
// }
