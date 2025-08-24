package consumer

import (
	"context"
	"user_service/global"
	"user_service/internal/eventbus/handler"
	"user_service/internal/eventbus/publisher"
)

func RunConsumer(ctx context.Context) {
	dlqPublisher := publisher.NewDLQPublisher()

	syncAuthDBConsumer := &SyncAuthDBConsumer{
		logger:       global.Logger,
		dlqPublisher: dlqPublisher,
		handler:      handler.NewSyncAuthHandler(),
	}

	syncAuthDBConsumer.ConsumeUserDB(ctx)
	global.Logger.Info("Sync Auth DB Consumer started", "")
}
