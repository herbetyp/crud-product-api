package middlewares

import (
	// "fmt"
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

		ParamUserId := c.Param("user_id")
		tokenString := authHeader[len(BearerScheme):]

		isValidToken, jwtSub := services.ValidateToken(tokenString, ParamUserId)
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

		AdminConf := configs.GetConfig().ADMIN

		ParamuserID, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "ID has to be integer"})
			return
		}

		SubUserID, _ := strconv.Atoi(c.Param("jwt_sub"))

		repo := &repositories.UserRepository{}
		user, err := repo.Get(uint(SubUserID))

		if err != nil {
			fmt.Println("Failed to get user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if user.UId == AdminConf.UId {
			c.Next()
			return

		} else {
			if ParamuserID == SubUserID {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
