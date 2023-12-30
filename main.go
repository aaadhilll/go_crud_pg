package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/controllers"
	"github.com/robbyklein/go-crud/initializers"
	// "gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// fmt.Println("Hello world")
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pongs",
	// 	})
	// })
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostDelete)

	// test russ

	r.POST("/twonums", controllers.PostsNums)
	r.Run() // listen and serve on 0.0.0.0:8080
}
