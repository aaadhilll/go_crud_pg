package main

import (
	// "fmt"
	"fmt"

	_ "github.com/robbyklein/go-crud/docs"
	"github.com/robbyklein/go-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/controllers"
	"github.com/robbyklein/go-crud/initializers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "gorm.io/gorm"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:3000
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.CountryCity{})
}

func main() {
	fmt.Println("Hello world main")
	r := gin.Default()
	/// swagger router adding
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	////
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

	// api calling and insert to table
	r.GET("/country/city", controllers.GetAllCity)

	// test 

	r.POST("/twonums", controllers.PostsNums)
	r.Run() // listen and serve on 0.0.0.0:8080
}
