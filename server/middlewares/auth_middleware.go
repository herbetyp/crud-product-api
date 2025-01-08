package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerScheme = "Bearer "

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Request does not contain an access token"})
			return
		}

		if len(authHeader) <= len(BearerScheme) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Invalid authorization header format"})
			return
		}
		userId := ctx.Param("user_id")

		tokenString := authHeader[len(BearerScheme):]
		if !services.ValidateToken(tokenString, userId) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
