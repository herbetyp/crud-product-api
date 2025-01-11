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

		products.GET("", controllers.GetProducts)
		products.GET("/:product_id", controllers.GetProduct)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:product_id", controllers.UpdateProduct)
		products.DELETE("/:product_id", controllers.DeleteProduct)

		users.GET("", controllers.GetUsers)
		users.GET("/:user_id", controllers.GetUser)
		users.POST("", controllers.CreateUser)
		users.POST("/login", controllers.Login)
		users.PATCH("/:user_id", middlewares.AuthMiddlewareUser(), controllers.UpdateUser)
		users.DELETE("/:user_id", middlewares.AuthMiddlewareAdmin(), controllers.DeleteUser)
	}

	return router
}
