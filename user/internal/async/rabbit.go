package async

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
	"github.com/wagslane/go-rabbitmq"
	internalLogger "log"
	"os"
)

type rabbit struct {
	consumer *rabbitmq.Consumer
	uc       *service.User
}

func newRabbit(targetAddr string) *rabbit {
	conn, err := rabbitmq.NewConn(
		fmt.Sprintf("amqp://guest:guest@%s", targetAddr),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		internalLogger.Println("failed to connect to RabbitMQ: ", err)
		return nil // Add proper handling if connection fails
	}
	internalLogger.Println("Consumer connected successfully")

	consumer, err := rabbitmq.NewConsumer(
		conn,
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey(os.Getenv("RABBITMQ_ROUTING_KEY")),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		internalLogger.Println("failed to initialize consumer: ", err)
		return nil // Add proper handling if consumer initialization fails
	}
	internalLogger.Println("Consumer initialized successfully")

	return &rabbit{consumer: consumer}
}

func (r *rabbit) Consume(ctx context.Context) <-chan []byte {
	ch := make(chan []byte, 1)

	go func() {
		err := r.consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
			fmt.Println("Message consumed")
			ch <- d.Body
			return rabbitmq.Ack
		})
		if err != nil {
			internalLogger.Println("error running consumer: ", err)
		}
	}()

	return ch
}
