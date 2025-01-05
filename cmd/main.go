package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/controllers"
	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/handlers"
	"github.com/herbetyp/crud-product-api/repository"
)

func main() {
	// Initialize the server
	server := gin.Default()

	// Grouping the API version
	v1 := server.Group("/v1")

	// Database connection
	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Repositories
	productRepository := repository.NewProductRepository(dbConnection)
	// Handlers
	productHandler := handlers.NewProductHandler(productRepository)
	// Controllers
	productController := controllers.NewProductController(productHandler)
	// Test the connection to server
	server.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
	// API Endpoints
	v1.GET("/products", productController.GetProducts)
	v1.GET("/product/:id", productController.GetProductById)
	v1.POST("/product", productController.CreateProduct)
	v1.PUT("/product/:id", productController.UpdateProduct)
	v1.DELETE("/product/:id", productController.DeleteProduct)

	// Run the server port
	cfg := configs.GetConfig()
	server.Run(fmt.Sprintf(":%s", cfg.API.Port))
}
