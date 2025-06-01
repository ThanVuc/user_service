package controller

import (
	"user_service/internal/services"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.IAuthService
}

func CreateAuthController(
	userService services.IAuthService,
) *AuthController {
	return &AuthController{
		authService: userService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login successful",
	})
}

func (ac *AuthController) ForgotPassword(c *gin.Context) {}

func (ac *AuthController) RefreshToken(c *gin.Context) {}

func (ac *AuthController) Logout(c *gin.Context) {}

func (ac *AuthController) CreateAuthor(c *gin.Context) {
	panic(response.BadRequest("CreateAuthor method is not implemented yet"))
}

func (ac *AuthController) CreateUserAndAuthor(c *gin.Context) {
	isCreated := ac.authService.CreateUserAndAuthor(c)
	if isCreated {
		response.Created(c, "Created user and author successfully", nil)
	} else {
		panic(response.InternalServerError("Failed to create user and author"))
	}
}
