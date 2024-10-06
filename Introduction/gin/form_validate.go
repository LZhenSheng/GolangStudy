package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // 绑定为json
// type LoginForm struct {
// 	User     string `form:"user" json:"user" xml:"user"  binding:"required,min=3,max=10"` //min最短长度，max最大长度
// 	Password string `form:"password" json:"password" xml:"password" binding:"required"`
// }

// func main() {
// 	router := gin.Default()
// 	router.POST("/loginJSON", func(c *gin.Context) {
// 		var loginForm LoginForm
// 		if err := c.ShouldBind(&loginForm); err != nil {
// 			fmt.Println(err.Error())
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"msg": "登录成功",
// 		})
// 	})
// 	router.Run()
// }
