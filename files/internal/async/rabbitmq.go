package async

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/Chatter/files/internal/log"
	"github.com/wagslane/go-rabbitmq"
	"os"
)

var _ Publisher = &publisher{}

type publisher struct {
	pub *rabbitmq.Publisher
}

func NewPublisher(ctx context.Context, targetAddr string) Publisher {
	conn, err := rabbitmq.NewConn(
		fmt.Sprintf("amqp://guest:guest@%s", targetAddr),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.GetLogger(ctx).Error("failed to connect to publisher: " + err.Error())
		return nil // Add proper handling if connection fails
	}
	log.GetLogger(ctx).Info("Publisher connected successfully")

	pub, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.GetLogger(ctx).Error("failed to init publisher: " + err.Error())
		return nil // Add proper handling if publisher initialization fails
	}
	log.GetLogger(ctx).Info("Publisher initialized successfully")

	return &publisher{pub: pub}
}

func (p *publisher) Publish(ctx context.Context, message any) {
	res, err := json.Marshal(message)
	if err != nil {
		log.GetLogger(ctx).Error("failed to marshal to publisher: " + err.Error())
		return
	}

	err = p.pub.Publish(
		res,
		[]string{os.Getenv("RABBITMQ_ROUTING_KEY")},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
	if err != nil {
		log.GetLogger(ctx).Error("failed to send to queue: " + err.Error())
		return
	}
	log.GetLogger(ctx).Info("Message published successfully")
}
