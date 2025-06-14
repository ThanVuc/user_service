package middlewares

import (
	grpc_client "user_service/internal/grpc/client"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

func CheckPerm(resource string, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticationClient, err := grpc_client.NewAuthenticationClient()
		if err != nil {
			panic(response.ErrorResponse{
				StatusCode: 403,
				Message:    "Forbidden",
				CodeReason: "Forbidden to access this resource",
			})
		}

		jwt := "asdkjsdfjskdfjg0"
		isValid, err := authenticationClient.ValidateToken(jwt)
		if err != nil || !isValid {
			panic(response.ErrorResponse{
				StatusCode: 403,
				Message:    "Forbidden",
				CodeReason: "Invalid token",
			})
		}
		println("Token is valid")

		// Check if the user has permission for the requested resource and action

		c.Next() // Proceed to the next middleware or handler if permission is granted
	}
}
