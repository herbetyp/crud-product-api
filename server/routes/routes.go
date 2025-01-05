package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1")

	{
		base_url.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
		base_url.GET("/products", controllers.GetProductsController)
		base_url.GET("/product/:id", controllers.GetProductByIdController)
		base_url.POST("/product", controllers.CreateProductController)
		base_url.PUT("/product/:id", controllers.UpdateProductController)
		base_url.DELETE("/product/:id", controllers.DeleteProductController)
	}
	return router
}
