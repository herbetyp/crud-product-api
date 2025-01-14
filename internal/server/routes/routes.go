package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/controllers"
	"github.com/herbetyp/crud-product-api/internal/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1", middlewares.RequestMiddleware())

	products := base_url.Group("/products", middlewares.AuthMiddleware())

	users := base_url.Group("/users", middlewares.AuthMiddleware())

	base_url.POST("/oauth2/token", controllers.Login)

	{
		base_url.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })

		products.GET("", controllers.GetProducts)
		products.GET("/:product_id", controllers.GetProduct)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:product_id", controllers.UpdateProduct)
		products.DELETE("/:product_id", controllers.DeleteProduct)

		users.GET("", middlewares.AuthMiddlewareAdmin(), controllers.GetUsers)
		users.GET("/:user_id", middlewares.AuthMiddlewareAdmin(), controllers.GetUser)
		users.POST("", controllers.CreateUser)
		users.PATCH("/:user_id/passwd-update", middlewares.AuthMiddlewareAdmin(), controllers.UpdateUser)
		users.DELETE("/:user_id", middlewares.AuthMiddlewareAdmin(), controllers.DeleteUser)
		users.POST("/:user_id/recovery", middlewares.AuthMiddlewareAdmin(), controllers.RecoveryUser)
	}

	return router
}
