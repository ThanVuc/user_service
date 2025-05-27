package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TrackLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("x-request-id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set("requestId", requestId)
		c.Next()
	}
}
