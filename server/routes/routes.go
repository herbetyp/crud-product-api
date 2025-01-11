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

		products.GET("", controllers.GetAll)
		products.GET("/:product_id", controllers.Get)
		products.POST("", controllers.Create)
		products.PUT("/:product_id", controllers.Update)
		products.DELETE("/:product_id", controllers.Delete)

		users.POST("", controllers.CreateUserContoller)
		users.POST("/login", controllers.LoginController)
		users.PATCH("/:user_id/reset-password", middlewares.AuthMiddlewareUser(), controllers.UpdateUserPassController)
		users.DELETE("/:user_id", middlewares.AuthMiddlewareAdmin(), controllers.DeleteUserController)
	}

	return router
}
