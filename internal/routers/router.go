package routers

import (
	"user_service/internal/controller"
	"user_service/internal/helper"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("v1/api")
	{
		user := v1.Group("user")
		{
			user.GET("/1", controller.CreateUserController().GetUserInfo)
			helper.RegisterRoute("GET", "/v1/api/user/1")
			user.GET("error-test", controller.CreateUserController().ErrorHandleTest)
			helper.RegisterRoute("GET", "/v1/api/user/error-test")
		}
	}
}
