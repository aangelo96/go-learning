package main

import (
	"github.com/aangelo96/go-learning/controllers"
	"github.com/aangelo96/go-learning/database"
	"github.com/aangelo96/go-learning/models"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectToDB()
}

func main() {
	// Auto migrate DB
	database.DB.AutoMigrate(&models.Product{})

	router := gin.Default()

	// Routes
	router.POST("/product", controllers.AddProduct)
	router.PUT("/product/:id", controllers.UpdateProductById)
	router.GET("/product/:id", controllers.GetProductById)
	router.GET("/product/all", controllers.GetProducts)
	router.DELETE("/product/:id", controllers.DeleteProductById)

	router.Run("localhost:8080")
}
