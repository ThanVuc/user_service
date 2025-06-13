package producers

import (
	"user_service/global"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func SendApiResourceCreatedEvent(serviceName string, resource []byte) error {
	logger := global.Logger
	channel := global.RabbitMQSharedChannel
	exchangeName := "direct_create_resource"
	routingKey := "api.resource.created"

	err := channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		logger.ErrorString("Failed to declare exchange", zap.String("exchange", exchangeName), zap.Error(err))
		return err
	}

	err = channel.Publish(
		exchangeName,
		routingKey,
		false, // mandatory
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        resource,
			Headers: map[string]interface{}{
				"service": serviceName,
			},
		},
	)

	if err != nil {
		logger.ErrorString("Failed to publish message", zap.String("exchange", exchangeName), zap.String("routingKey", routingKey), zap.Error(err))
		return err
	}

	logger.InfoString("Resource created event sent",
		zap.String("exchange", exchangeName),
		zap.String("routingKey", routingKey),
		zap.ByteString("resource", resource))

	return nil
}
