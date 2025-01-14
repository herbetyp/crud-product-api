package middlewares

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	config "github.com/herbetyp/crud-product-api/internal/configs"
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
			return
		}

		if len(authHeader) <= len(BearerScheme) {
			c.AbortWithStatusJSON(401,
				gin.H{"error": "Invalid authorization header format"})
			return
		}

		ParamUserId := c.Param("user_id")
		tokenString := authHeader[len(BearerScheme):]

		ok, _ := services.ValidateToken(tokenString, ParamUserId)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		AdminConf := config.GetConfig().ADMIN

		ParamUserID, err := strconv.Atoi(c.Param("user_id"))

		if ParamUserID > 0 && err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "ID has to be integer"})
			return
		}

		const BearerScheme = "Bearer "

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerScheme):]

		token, _, _ := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})

		claims, _ := token.Claims.(jwt.MapClaims)

		SubUserID, _ := strconv.Atoi(claims["sub"].(string))

		userUID, err := repositories.GetUID(uint(SubUserID))

		if err != nil {
			fmt.Println("Failed to get user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if userUID == AdminConf.UId {
			c.Next()
			return
		} else {
			if ParamUserID == SubUserID {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
