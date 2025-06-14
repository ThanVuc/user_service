package middlewares

import (
	grpc_client "user_service/internal/grpc/client"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

func CheckPerm(resource string, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := "asdkjsdfjskdfjg0"
		authenticationClient, err := grpc_client.NewAuthenticationClient()
		if err != nil {
			panic(response.ErrorResponse{
				StatusCode: 403,
				Message:    "Forbidden",
				CodeReason: "Forbidden to access this resource",
			})
		}
		isValid, err := authenticationClient.ValidateToken(jwt)
		if err != nil || !isValid {
			panic(response.ErrorResponse{
				StatusCode: 403,
				Message:    "Forbidden",
				CodeReason: "Invalid token",
			})
		}
		println("Token is valid")

		authorizationClient, err := grpc_client.NewAuthorizationClient()
		if err != nil {
			panic(response.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
				CodeReason: "Internal Server Error",
			})
		}

		allowed, err := authorizationClient.CheckPerm(resource, action, jwt)
		if err != nil || !allowed {
			panic(response.ErrorResponse{
				StatusCode: 403,
				Message:    "Forbidden",
				CodeReason: "You do not have permission to access this resource",
			})
		}

		println("Permission granted for resource:", resource, "and action:", action)

		c.Next() // Proceed to the next middleware or handler if permission is granted
	}
}
