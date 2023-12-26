package main

import (
	"github.com/robbyklein/go-crud/initializers"
	"github.com/robbyklein/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}