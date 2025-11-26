package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kasahun-welela/go-movie-webserver/config"
	"github.com/kasahun-welela/go-movie-webserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetMovies fetches all movies
func GetMovies(c *gin.Context) {
	collection := config.GetCollection("movies")

	// Context with timeout for DB ops to avoid hanging requests.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all documents (empty filter).
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching movies"})
		return
	}
	// Ensure cursor is closed once done.
	defer cursor.Close(ctx)

	var movies []models.Movie
	if err := cursor.All(ctx, &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}
