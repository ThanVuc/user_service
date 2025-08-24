package initialize

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"user_service/global"
	"user_service/internal/eventbus/consumer"

	"go.uber.org/zap"
)

func Run() {
	if err := initConfigAndResources(); err != nil {
		global.Logger.Error("Failed to initialize configs and resources", "", zap.Error(err))
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	startGrpcSerivces(ctx, wg)
	consumer.RunConsumer(ctx)

	<-stop
	global.Logger.Info("Shutdown signal received, shutting down...", "")

	cancel()

	gracefulShutdown(wg, global.Logger)
}
