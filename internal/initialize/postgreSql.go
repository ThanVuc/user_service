package initialize

import (
	"fmt"
	"time"
	"user_service/global"
	po "user_service/internal/PO"
	logstruct "user_service/internal/log_struct"
	"user_service/pkg/response"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func checkInitError(err error) {
	if err != nil {
		panic("failed to initialize PostgreSQL: " + err.Error())
	}
}

func InitPostgreSQL() {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	configs := global.Config.Postgres
	var connectString = fmt.Sprintf(dsn, configs.Host, configs.User, configs.Password, configs.Database, configs.Port)
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkInitError(err)
	global.PostgresDb = db

	// set pool
	setPool()
	migrateTables()
	go trackConnection()
}

func setPool() {
	postConfig := global.Config.Postgres
	postgresDb, err := global.PostgresDb.DB()
	checkInitError(err)
	postgresDb.SetMaxIdleConns(postConfig.MaxIdleConns)
	postgresDb.SetConnMaxIdleTime(time.Duration(postConfig.ConnMaxIdleTime) * time.Minute) // in seconds
	postgresDb.SetMaxOpenConns(postConfig.MaxOpenConns)
	postgresDb.SetConnMaxLifetime(time.Duration(postConfig.MaxLifetime) * time.Second.Abs()) // in seconds
}

func migrateTables() {
	err := global.PostgresDb.AutoMigrate(
		&po.User{},
	)

	checkInitError(err)
	logger := global.Logger
	logger.Info(logstruct.NewAuthenLog("nil", "Migrate Table to postgre successful!"))
}

func trackConnection() {
	postgresDb, err := global.PostgresDb.DB()
	logger := global.Logger
	if err != nil {
		fmt.Println("failed to get PostgreSQL DB instance:", err)
		return
	}

	statTicker := time.NewTicker(5 * time.Second)
	defer statTicker.Stop()

	for {
		select {
		case <-statTicker.C:
			stats := postgresDb.Stats()
			if stats.OpenConnections > stats.MaxOpenConnections {
				logger.Error(response.InternalServerError("Overload the connections"), "routine check", nil)
			}
		}
	}
}
