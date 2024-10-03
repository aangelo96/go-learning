package controllers

import (
	"github.com/aangelo96/go-learning/database"
	"github.com/aangelo96/go-learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create a new product
func AddProduct(c *gin.Context) {

	var body struct {
		Name        string
		Description string
		Price       float64
		Stock       int
		Status      int
		WebId       uint
		Image       string
		Order       int
		Group       string
	}

	c.Bind(&body)

	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		Status:      body.Status,
		WebId:       body.WebId,
		Image:       body.Image,
		Order:       body.Order,
		Group:       body.Group,
	}

	database.DB.Create(&product)
	c.IndentedJSON(http.StatusOK, product)
}

// Get Product by ID
func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var prod models.Product
	err := database.DB.First(&prod, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, prod)
}

// Get Products
func GetProducts(c *gin.Context) {
	var products []models.Product
	err := database.DB.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

// Update Product by ID
func UpdateProductById(c *gin.Context) {
	var body struct {
		Name        string
		Description string
		Price       float64
		Stock       int
		Status      int
		WebId       uint
		Image       string
		Order       int
		Group       string
	}

	c.Bind(&body)
	id := c.Param("id")

	var prod models.Product
	err := database.DB.First(&prod, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	prod.Name = body.Name
	prod.Description = body.Description
	prod.Price = body.Price
	prod.Stock = body.Stock
	prod.Status = body.Status
	prod.WebId = body.WebId
	prod.Order = body.Order
	prod.Group = body.Group
	database.DB.Save(&prod)

	c.IndentedJSON(http.StatusOK, prod)
}

// Get Product by ID
func DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	var prod models.Product
	err := database.DB.First(&prod, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	database.DB.Delete(&prod)
	c.IndentedJSON(http.StatusOK, prod)
}
