package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1")

	products := base_url.Group("/products")
	users := base_url.Group("/users")
	{
		base_url.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })

		products.GET("", controllers.GetProductsController)
		products.GET("/:id", controllers.GetProductByIdController)
		products.POST("", controllers.CreateProductController)
		products.PUT("/:id", controllers.UpdateProductController)
		products.DELETE("/:id", controllers.DeleteProductController)

		users.POST("/created", controllers.CreateUserContoller)
		users.POST("/login", controllers.LoginController)
	}
	return router
}
