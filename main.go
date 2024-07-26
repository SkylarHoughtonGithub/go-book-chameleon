package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Redirect from root path to /static/index.html
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/index.html")
	})

	// API routes
	api := router.Group("/api")
	{
		// Book routes
		api.GET("/books", getBooks)
		api.GET("/books/:index", getBook)
		api.POST("/books", addBook)
		api.PUT("/books/:index", updateBook)
		api.DELETE("/books/:index", deleteBook)

		// Animal routes
		api.GET("/animals", getAnimals)
		api.GET("/animals/:index", getAnimal)
		api.POST("/animals", addAnimal)
		api.PUT("/animals/:index", updateAnimal)
		api.DELETE("/animals/:index", deleteAnimal)
	}

	// Start server
	router.Run(":8080")
}
