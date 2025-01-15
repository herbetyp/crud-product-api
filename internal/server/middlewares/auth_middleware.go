package middlewares

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/herbetyp/crud-product-api/repositories"
	"github.com/herbetyp/crud-product-api/services"
	zapLog "go.uber.org/zap"
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

		tokenString := authHeader[len(BearerScheme):]

		ok, tokenId, subId, _ := services.ValidateToken(tokenString)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		subIdInt, _ := strconv.Atoi(subId)
		user, _ := repositories.GetInfo(uint(subIdInt))
		// Log
		logger.Info("Authorized",
			zapLog.String("request_id", c.Param("X-Request-Id")),
			zapLog.String("ip", c.Param("ip")),
			zapLog.String("username", user.Username),
			zapLog.String("email", user.Email),
			zapLog.String("jwt_id", tokenId),
			zapLog.String("method", c.Param("method")),
			zapLog.String("path", c.Param("path")),
			zapLog.String("protocol", c.Param("protocol")),
			zapLog.String("user_agent", c.Param("user_agent")),
		)
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

		user, err := repositories.GetInfo(uint(SubUserID))

		if err != nil {
			fmt.Println("Failed to get user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if user.UId == AdminConf.UId {
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
