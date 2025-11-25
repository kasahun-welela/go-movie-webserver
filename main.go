package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Welcome to movie web-server"})
	})
	router.Run("localhost:8080")
	fmt.Println("Hello world")
}
