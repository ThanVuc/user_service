package controller

import (
	"user_service/internal/services"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		response.NotFound(c, "users not found")
		return
	}

	c.JSON(200, gin.H{"users": users})
}

func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	response.Ok(c, "update user info successfully", nil)
}
