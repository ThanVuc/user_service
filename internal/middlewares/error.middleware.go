package middlewares

import (
	"net/http"
	"runtime/debug"
	"user_service/global"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger := global.Logger
				requestIdVal, _ := c.Get("requestId")
				requestId, _ := requestIdVal.(string)
				switch e := r.(type) {
				case response.ErrorResponse:
					logger.Error(e, requestId, debug.Stack())
					c.JSON(e.StatusCode, e)
				default:
					logger.Error(response.InternalServerError("Unknown panic"), requestId, debug.Stack())
					c.JSON(500, response.AnotherError(http.StatusInternalServerError, "Unknown panic"))
				}
			}

			c.Abort()
		}()
		c.Next()
	}
}
