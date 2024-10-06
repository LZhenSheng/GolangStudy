package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Person struct {
// 	ID   string `uri:"id" binding:"required,uuid"`
// 	Name string `uri:"name" binding:"required"`
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/:name/:id", func(c *gin.Context) {
// 		var person Person
// 		if err := c.ShouldBindUri(&person); err != nil {
// 			c.Status(404)
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"name": person.Name,
// 			"id":   person.ID,
// 		})
// 	})
// 	router.Run() // listen and serve on 0.0.0.0:8080
// 	// r.Run(":8083") // listen and serve on 0.0.0.0:8080
// }
