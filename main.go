package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kasahun-welela/go-movie-webserver/config"
	controller "github.com/kasahun-welela/go-movie-webserver/controllers"
)

func main() {
	config.ConnectDB()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Welcome to movie web-server"})
	})
	router.GET("/movies", controller.GetMovies)
	router.Run("localhost:8080")
	fmt.Println("Hello world")
}
