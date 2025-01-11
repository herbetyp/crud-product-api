package middlewares

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/repositories"
	"github.com/herbetyp/crud-product-api/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerScheme = "Bearer "

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401,
				gin.H{"error": "Request does not contain an access token"})
			return
		}

		if len(authHeader) <= len(BearerScheme) {
			ctx.AbortWithStatusJSON(401,
				gin.H{"error": "Invalid authorization header format"})
			return
		}
		userId := ctx.Param("user_id")

		tokenString := authHeader[len(BearerScheme):]
		isValidToken, jwtSub := services.ValidateToken(tokenString, userId)
		if !isValidToken {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.AddParam("jwt_sub", jwtSub)
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(ctx)

		conf := configs.GetConfig()
		
		userID, err := strconv.Atoi(ctx.Param("jwt_sub"))
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "ID has to be integer"})
			return
		}
		
		repo := &repositories.UserRepository{}	
		user, err := repo.Get(uint(userID))

		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Failed to get user"})
			return
		}
		
		if user.UId != conf.ADMIN.UId {
			ctx.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}

		if ctx.Param("user_id") == ctx.Param("jwt_sub") {
			ctx.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}
	}}
func AuthMiddlewareUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(ctx)

		subClaim, _ := strconv.Atoi(ctx.Param("jwt_sub"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))

		if subClaim != userId {
			fmt.Printf("invalid user: uid not match\n")
			ctx.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
