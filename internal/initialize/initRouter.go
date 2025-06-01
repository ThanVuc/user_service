package initialize

import (
	"user_service/internal/helper"
	"user_service/internal/middlewares"
	"user_service/internal/routers"
	"user_service/internal/wire"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.TrackLogMiddleware())
	// cor
	// limiter global
	authRouter := routers.RouterGroupApp.AuthRouterEnter.AuthRouter
	authController := wire.InjectAuthController()

	MainGroup := r.Group("v1/api")
	{
		MainGroup.GET("/checkStatus", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Authentication Service is running",
			})
		})

		MainGroup.POST("/test/createAuthor", authController.CreateAuthor)               // Create an author
		MainGroup.POST("/test/createUserAndAuthor", authController.CreateUserAndAuthor) // Create a user and author
	}
	{
		authRouter.InitAuthRouter(MainGroup) // Initialize the auth router group
	}

	helper.WriteRouteToFile()
}
