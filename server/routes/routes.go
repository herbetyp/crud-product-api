package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
	"github.com/herbetyp/crud-product-api/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1")

	products := base_url.Group("/products", middlewares.AuthMiddleware())
	users := base_url.Group("/users")
	{
		base_url.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })

		products.GET("", controllers.GetProductsController)
		products.GET("/:product_id", controllers.GetProductByIdController)
		products.POST("", controllers.CreateProductController)
		products.PUT("/:product_id", controllers.UpdateProductController)
		products.DELETE("/:product_id", controllers.DeleteProductController)

		users.POST("", controllers.CreateUserContoller)
		users.POST("/login", controllers.LoginController)
		users.PATCH("/:user_id/reset-password", middlewares.AuthMiddleware(), controllers.UpdateUserPassController)
		users.DELETE("/:user_id", middlewares.AuthMiddleware(), controllers.DeleteUserController)
	}

	return router
}
