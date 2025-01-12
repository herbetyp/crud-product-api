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
	return func(c *gin.Context) {
		const BearerScheme = "Bearer "

		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Authorization header is required"})
			c.AddParam("auth", "unauthorized")
			return
		}

		if len(authHeader) <= len(BearerScheme) {
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Invalid authorization header format"})
				c.AddParam("auth", "unauthorized")
				return
		}

		userId := c.Param("user_id")
		tokenString := authHeader[len(BearerScheme):]

		isValidToken, jwtSub := services.ValidateToken(tokenString, userId)
		if !isValidToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			c.AddParam("auth", "unauthorized")
			return
		}

		c.AddParam("jwt_sub", jwtSub)
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(c)

		if c.Param("auth") == "unauthorized" {
			c.Abort()
			return
		}
		
		conf := configs.GetConfig()

		userID, err := strconv.Atoi(c.Param("jwt_sub"))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "ID has to be integer"})
			return
		}

		repo := &repositories.UserRepository{}
		user, err := repo.Get(uint(userID))

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to get user"})
			return
		}

		if user.UId != conf.ADMIN.UId {
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}

		if c.Param("user_id") == c.Param("jwt_sub") {
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
func AuthMiddlewareUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware := AuthMiddleware()
		authMiddleware(c)

		if c.Param("auth") == "unauthorized" {
			c.Abort()
			return
		}

		subClaim, _ := strconv.Atoi(c.Param("jwt_sub"))
		userId, _ := strconv.Atoi(c.Param("user_id"))

		if subClaim != userId {
			fmt.Printf("invalid user: uid not match\n")
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
