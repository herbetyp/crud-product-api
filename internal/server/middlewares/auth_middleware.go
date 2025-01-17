package middlewares

import (
	"encoding/json"
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

		var u userModel.User

		USER_PREFIX := "user>>>"
		cacheKey := USER_PREFIX + claims["fingerprint"].(string)

		cacheData := services.GetCache(cacheKey)
		if cacheData != "" {
			err := json.Unmarshal([]byte(cacheData), &u)
			if err != nil {
				return
			}
		} else {
			repo := &repository.UserRepository{}
			handler := handlers.NewUserHandler(repo)

			sub, _ := strconv.ParseUint(claims["sub"].(string), 10, 32)
			u, _ = handler.GetUser(uint(sub), false)

			cacheValue, _ := json.Marshal(u)
			services.SetCache(cacheKey, string(cacheValue))
		}

		// Log
		logger.Info("Authorized",
			zapLog.String("request_id", c.Param("X-Request-Id")),
			zapLog.String("ip", c.Param("ip")),
			zapLog.String("username", u.Username),
			zapLog.String("email", u.Email),
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

		// if ParamUserID > 0 && err != nil {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "ID has to be integer"})
		// 	return
		// }

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

		cacheData := services.GetCache(cacheKey)
		if cacheData != "" {
			err := json.Unmarshal([]byte(cacheData), &u)
			if err != nil {
				return
			}
		} else {
			repo := &repository.UserRepository{}
			handler := handlers.NewUserHandler(repo)

			u, _ := handler.GetUser((uint(claims["sub"].(float64))), false)

			cacheValue, _ := json.Marshal(u)
			services.SetCache(cacheKey, string(cacheValue))
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
