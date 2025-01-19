package middlewares

import (
	// "encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	userModel "github.com/herbetyp/crud-product-api/models/user"
	repository "github.com/herbetyp/crud-product-api/repositories"
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

		ok, claims, err := services.ValidateToken(tokenString)

		if !ok || err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		var user userModel.User

		USER_PREFIX := "user>>>"
		cacheKey := USER_PREFIX + claims["fingerprint"].(string)

		repo := &repository.UserRepository{}
		handler := handlers.NewUserHandler(repo)

		cached := services.GetCache(cacheKey, &user)
		if cached == "" {
			sub, _ := strconv.ParseUint(claims["sub"].(string), 10, 32)
			u, err := handler.GetUser(uint(sub), false)
			if err != nil {
				logger.Error("error on get user from database: %v", err)
				return
			}
			services.SetCache(cacheKey, &u)
		}

		// Log
		logger.Info("Authorized",
			zapLog.String("request_id", c.Param("X-Request-Id")),
			zapLog.String("ip", c.Param("ip")),
			zapLog.String("username", user.Username),
			zapLog.String("email", user.Email),
			zapLog.String("jwt_id", claims["jti"].(string)),
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

		ParamUserID, _ := strconv.Atoi(c.Param("user_id"))

		const BearerScheme = "Bearer "

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerScheme):]

		claims, err := services.GetJwtClaims(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		var u userModel.User

		USER_PREFIX := "user>>>"
		cacheKey := USER_PREFIX + claims["fingerprint"].(string)

		repo := &repository.UserRepository{}
		handler := handlers.NewUserHandler(repo)

		sub, _ := strconv.ParseUint(claims["sub"].(string), 10, 32)
		u, _ = handler.GetUser(uint(sub), false)

		cached := services.GetCache(cacheKey, &u)
		if cached == "" {
			u, err := handler.GetUser(uint(sub), false)
			if err != nil {
				logger.Error("error on get user from database: %v", err)
				return
			}
			services.SetCache(cacheKey, &u)
		}

		if u.UId == AdminConf.UId {
			c.Next()
			return
		} else {
			if ParamUserID == int(u.ID) {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
