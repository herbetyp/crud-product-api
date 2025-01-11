package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/repositories"
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
		isValidToken, subClaim := services.ValidateToken(tokenString, userId)
		if !isValidToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.AddParam("sub", subClaim)
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(ctx)

		conf := configs.GetConfig()

		userID, err := strconv.Atoi(ctx.Param("sub"))
		if err != nil {
			response := models.Response{Message: "Invalid user ID. Only integer are allowed"}
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		user, _ := repositories.GetUserByIdRepository(userID)
		if user.UId != conf.ADMIN.UId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
		
		if ctx.Param("user_id") == ctx.Param("sub") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
func AuthMiddlewareUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(ctx)

		subClaim, _ := strconv.Atoi(ctx.Param("sub"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))

		if subClaim != userId {
			fmt.Printf("invalid user: uid not match\n")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
