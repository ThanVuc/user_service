package initialize

import (
	"os"
	"user_service/global"

	"github.com/thanvuc/go-core-lib/log"
)

func InitLogger() {
	env := os.Getenv("GO_ENV")
	global.Logger = log.NewLogger(log.Config{
		Env:   env,
		Level: global.Config.Log.Level,
	})
}
