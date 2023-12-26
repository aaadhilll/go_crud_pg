package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/initializers"
	"github.com/robbyklein/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// data of body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}

	// // Return it

	c.JSON(200, gin.H{
		"post": post,
	})

	// c.JSON(200, gin.H{
	// 	"post" : "helos",
	// })
}

func PostsIndex(c *gin.Context) {
	// get all post
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id of url
	id := c.Param("id")

	// get  post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id of url
	id := c.Param("id")
	// Get data of request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find post where updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	// get id of url
	id := c.Param("id")

	// delete post
	initializers.DB.Delete(&models.Post{}, id)

	// respond

	c.Status(200)
}