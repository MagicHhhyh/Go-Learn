package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	type a struct {
		Name string `json:"name"`
		Age  int    `form:"age"`
	}
	b := a{
		Name: "testJson",
		Age:  1,
	}
	d := a{
		Name: "testD",
		Age:  2,
	}
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, d)
		c.JSON(http.StatusOK, b)
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
	})

	router.Run(":8000")
}
