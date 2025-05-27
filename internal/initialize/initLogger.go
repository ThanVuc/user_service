package initialize

import (
	"user_service/global"
	"user_service/pkg/loggers"
)

func InitLogger() {
	global.Logger = loggers.NewLogger(
		global.Config.Log,
	)
}
