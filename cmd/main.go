package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
	"github.com/herbetyp/crud-product-api/db"
	"github.com/herbetyp/crud-product-api/repository"
	"github.com/herbetyp/crud-product-api/usecase"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")

	// Database
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
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// API Endpoints
	v1.GET("/products", productController.GetProducts)
	v1.GET("/product/:id", productController.GetProductById)
	v1.POST("/product", productController.CreateProduct)
	server.Run(":3000")
}
