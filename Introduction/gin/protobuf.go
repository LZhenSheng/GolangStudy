package main

// import (
// 	"net/http"

// 	"GolangStudy/Introduction/proto"

// 	"github.com/gin-gonic/gin"
// )

// func moreJSON(c *gin.Context) {
// 	var msg struct {
// 		Name    string `json:"user"`
// 		Message string
// 		Number  int
// 	}
// 	msg.Name = "bobby"
// 	msg.Message = "测试json"
// 	msg.Number = 20
// 	c.JSON(http.StatusOK, msg)
// }
// func returnProto(c *gin.Context) {
// 	user := &proto.Teacher{
// 		Name:   "bobby",
// 		Course: []string{"python", "go", "java"},
// 	}
// 	c.ProtoBuf(http.StatusOK, user)
// }
// func main() {
// 	router := gin.Default()
// 	router.GET("/moreJSON", moreJSON)
// 	router.GET("/someProtoBuf", returnProto)
// 	router.GET("/pureJSON", func(c *gin.Context) {
// 		c.PureJSON(200, gin.H{
// 			"html": "<b>Hello, world!</b>",
// 		})
// 	})
// 	router.Run()
// }
