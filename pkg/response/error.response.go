package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This package provides a standardized way to handle error responses in the application.
*/

// ErrorResponse is a struct that represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	CodeReason string `json:"errorReason"`
	CreatedAt  string `json:"createdAt"`
}

// The method to return the error response in controller
func BadRequest(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(BAD_REQUEST),
		Message:    message,
		CodeReason: MSG[BAD_REQUEST],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func Unauthorized(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(UNAUTHORIZED),
		Message:    message,
		CodeReason: MSG[UNAUTHORIZED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func Forbidden(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(FORBIDDEN),
		Message:    message,
		CodeReason: MSG[FORBIDDEN],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func NotFound(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(NOT_FOUND),
		Message:    message,
		CodeReason: MSG[NOT_FOUND],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func MethodNotAllowed(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(METHOD_NOT_ALLOWED),
		Message:    message,
		CodeReason: MSG[METHOD_NOT_ALLOWED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func NotAcceptable(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(NOT_ACCEPTABLE),
		Message:    message,
		CodeReason: MSG[NOT_ACCEPTABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func Conflict(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(CONFLICT),
		Message:    message,
		CodeReason: MSG[CONFLICT],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func UnsupportedMediaType(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(UNSUPPORTED_MEDIA_TYPE),
		Message:    message,
		CodeReason: MSG[UNSUPPORTED_MEDIA_TYPE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func InternalServerError(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(INTERNAL_SERVER_ERROR),
		Message:    message,
		CodeReason: MSG[INTERNAL_SERVER_ERROR],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func ServiceUnavailable(c *gin.Context, message string) {
	response := ErrorResponse{
		StatusCode: int(SERVICE_UNAVAILABLE),
		Message:    message,
		CodeReason: MSG[SERVICE_UNAVAILABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}

func AnotherError(c *gin.Context, statusCode int, message string) {
	response := ErrorResponse{
		StatusCode: int(statusCode),
		Message:    message,
		CodeReason: "UNKNOWN_ERROR",
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	c.JSON(response.StatusCode, response)
}
