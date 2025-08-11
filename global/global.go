package global

import (
	"user_service/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thanvuc/go-core-lib/cache"
	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: This package defines global variables that are used throughout the application.
*/
var (
	Config            settings.Config
	Logger            log.Logger
	PostgresPool      *pgxpool.Pool
	RedisDb           *cache.RedisCache
	EventBusConnector *eventbus.RabbitMQConnector
)
