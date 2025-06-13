package initialize

import (
	"user_service/internal/amqp/producers"
	"user_service/internal/helper"
)

func initDefaultProducers() {
	go producers.SendApiResourceCreatedEvent("user_service", helper.GetResources())
}
