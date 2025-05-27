package global

import (
	"user_service/pkg/loggers"
	"user_service/pkg/settings"

	"gorm.io/gorm"
)

var (
	Config     settings.Config
	Logger     *loggers.LoggerZap
	PostgresDb *gorm.DB
)
