package initialize

import (
	"fmt"
	"strings"
	"user_service/global"
	"user_service/internal/helper"
	"user_service/internal/middlewares"
	"user_service/internal/routers"

	"github.com/gin-gonic/gin"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitRedis()
	// InitPostgreSQL()
	// init the app
	r := gin.New()
	r.Use(middlewares.TrackLogMiddleware())
	r.Use(middlewares.ErrorHandler())
	routers.InitRouter(r)

	fmt.Println(strings.Join(helper.GetRoutes(), "\n"))
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port)) // listen and serve on
}
