package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
	"github.com/herbetyp/crud-product-api/db"
	"github.com/herbetyp/crud-product-api/repository"
	"github.com/herbetyp/crud-product-api/usecase"
)

func main() {
	// Initialize the server
	server := gin.Default()

	// Grouping the API version
	v1 := server.Group("/v1")

	// Database connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Repositories
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Usecases
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// Controllers
	productController := controllers.NewProductController(ProductUsecase)
	// Test the connection to server
	server.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
	// API Endpoints
	v1.GET("/products", productController.GetProducts)
	v1.GET("/product/:id", productController.GetProductById)
	v1.POST("/product", productController.CreateProduct)
	v1.PUT("/product/:id", productController.UpdateProduct)
	v1.DELETE("/product/:id", productController.DeleteProduct)
	// Run the server port
	server.Run(":3000")
}
