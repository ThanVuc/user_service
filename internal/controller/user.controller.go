package controller

import (
	"user_service/global"
	logstruct "user_service/internal/log_struct"
	"user_service/internal/services"
	"user_service/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func CreateUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	// Simulate getting user info from a database
	logger := global.Logger
	logger.Info(logstruct.NewAuthenLog("12345", "GetUserInfo"))
	response.Ok(c, "User info retrieved successfully", uc.userService.GetUserInfo())
}

func (uc *UserController) ErrorHandleTest(c *gin.Context) {
	panic(response.InternalServerError("This is a test error"))
}
