package producer

import (
	"context"
	"encoding/json"
	"github.com/MaksKazantsev/Chatter/messages/internal/broker"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

var _ broker.Producer = &AsyncProducer{}

type AsyncProducer struct {
	*kafka.Writer
	topic string
}

func NewProducer(addr string, topic string) *AsyncProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
	}

	return &AsyncProducer{writer, topic}
}

func (a *AsyncProducer) Produce(ctx context.Context, message any) {
	b, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	err = a.WriteMessages(ctx, kafka.Message{Value: b})
}
