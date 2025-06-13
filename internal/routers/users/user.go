package users

import (
	"user_service/internal/helper"
	"user_service/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	userController := wire.InjectUserController()

	privateRouter := Router.Group("users")
	{
		privateRouter.GET("/", userController.GetUserById)
		privateRouter.PUT("/", userController.UpdateUserInfo)
	}

	ur.registerUser()
}

func (ar *UserRouter) registerUser() {
	helper.AddResource("userInfo", []string{
		"getOne",
		"Update",
	})
}
