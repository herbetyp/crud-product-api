package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AddParam("request_id", c.GetHeader("X-Request-Id"))
		c.AddParam("ip", c.ClientIP())
		c.AddParam("method", c.Request.Method)
		c.AddParam("path", c.Request.URL.Path)
		c.AddParam("protocol", c.Request.Proto)
		c.AddParam("user_agent", c.GetHeader("User-Agent"))
		c.Next()
	}
}
