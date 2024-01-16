package controllers

import (
	// "fmt"

	"io"

	"github.com/gin-gonic/gin"
	// "github.com/robbyklein/go-crud/initializers"
	// "github.com/robbyklein/go-crud/models"
	"fmt"
	"net/http"
	"strings"
)




func GetAllCity(c *gin.Context) {
	// get all post
	// var posts []models.Post
	// initializers.DB.Find(&posts)






  url := "https://countriesnow.space/api/v0.1/countries/cities"
  method := "POST"

  payload := strings.NewReader(`{
    "country": "India"

}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))


	c.JSON(200, gin.H{
		"city": "working",
	})
}