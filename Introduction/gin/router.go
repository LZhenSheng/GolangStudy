package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func goodList(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "goodList",
// 	})
// }
// func goodsDetail(c *gin.Context) {
// 	id := c.Param("id")
// 	action := c.Param("action")
// 	c.JSON(200, gin.H{
// 		"message": "goodsDetail",
// 		"id":      id,
// 		"action":  action,
// 	})
// }
// func createGoods(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "createGoods",
// 	})
// }
// func main() {
// 	router := gin.Default()
// 	goodsGroup := router.Group("/goods")
// 	{
// 		goodsGroup.GET("", goodList)
// 		goodsGroup.GET("/:id/:action", goodsDetail)
// 		// goodsGroup.GET("/:id/*action", goodsDetail)
// 		goodsGroup.POST("", createGoods)
// 	}

// 	router.Run() // listen and serve on 0.0.0.0:8080
// 	// r.Run(":8083") // listen and serve on 0.0.0.0:8080
// }
