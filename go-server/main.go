package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []product{
	{ID: "1", Name: "Item A", Price: 50.00},
	{ID: "2", Name: "Item B", Price: 50.50},
	{ID: "3", Name: "Item C", Price: 1.25},
}

// getProducts responds with the list of all albums as JSON.
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:id", getProductByID)
	router.POST("/products", postProducts)

	router.Run("localhost:8080")
}
