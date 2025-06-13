package middlewares

import (
	"net/http"
	"runtime/debug"
	"user_service/global"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

/*
@Author: Sinh
@Date:
@Description: Global error handler middleware for the application.
*/
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger := global.Logger
				requestIdVal, _ := c.Get("requestId")
				requestId, _ := requestIdVal.(string)
				switch e := r.(type) {
				case response.ErrorResponse:
					// Needing more details ? -> add debug.Stack() to the logger
					if e.StatusCode >= 500 {
						logger.Error(e, requestId, debug.Stack())
					} else {
						logger.Error(e, requestId, nil)
					}

					c.JSON(e.StatusCode, e)
				default:
					logger.Error(response.ErrorResponse{
						StatusCode: http.StatusInternalServerError,
						Message:    "An unexpected error occurred. Please try again later.",
						CodeReason: "INTERNAL_SERVER_ERROR",
					}, requestId, debug.Stack())
					response.AnotherError(c, 500, "An unexpected error occurred. Please try again later.")
				}
			}

			c.Abort()
		}()
		c.Next()
	}
}
