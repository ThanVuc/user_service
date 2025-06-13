package initialize

import (
	"user_service/internal/middlewares"
	"user_service/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.TrackLogMiddleware())
	// cor
	// limiter global
	userRouter := routers.RouterGroupApp.UserRouterEnter.UserRouter

	MainGroup := r.Group("v1/api")
	{
		MainGroup.GET("/checkStatus", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Authentication Service is running",
			})
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
}
