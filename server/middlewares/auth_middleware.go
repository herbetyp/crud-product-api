package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/services"
)


func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerScheme = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Request does not contain an access token"})
			return
		}

		tokenString := authHeader[len(BearerScheme):]
		if !services.NewJWTService().ValidateToken(tokenString) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}