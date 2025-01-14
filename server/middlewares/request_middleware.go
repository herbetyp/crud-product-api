package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/configs/logger"
	zapLog "go.uber.org/zap"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("request started",
			zapLog.String("request_id", c.GetHeader("X-Request-Id")),
			zapLog.String("ip", c.ClientIP()),
			zapLog.String("method", c.Request.Method),
			zapLog.String("path", c.Request.URL.Path),
			zapLog.String("protocol", c.Request.Proto),
			zapLog.String("user_agent", c.GetHeader("User-Agent")),
		)
		zapLog.AddCaller()
		c.Next()
	}
}
