package initialize

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"user_service/global"

	"go.uber.org/zap"
)

func Run() {
	if err := initConfigAndResources(); err != nil {
		global.Logger.Error("Failed to initialize configs and resources", "", zap.Error(err))
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	startGrpcSerivces(ctx, wg)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	global.Logger.Info("Shutdown signal received, shutting down...", "")

	cancel()

	gracefulShutdown(wg, global.Logger)
}
