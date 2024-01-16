package main

import (
	"fmt"
	

	"github.com/robbyklein/go-crud/initializers"
	"github.com/robbyklein/go-crud/models"
)

func init() {
	fmt.Println("migrate init")
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("migrate main")
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.CountryCity{})
}
