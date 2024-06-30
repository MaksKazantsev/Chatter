package broker

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/MaksKazantsev/Chatter/messages/internal/broker"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"log/slog"
)

var _ broker.Producer = &AsyncProducer{}

type AsyncProducer struct {
	sarama.AsyncProducer
	topic string
}

func NewProducer(addr string, topic string) *AsyncProducer {
	cfg := initKafkaConfig()

	prod, err := sarama.NewAsyncProducer([]string{addr}, cfg)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	defer func() {
		if err = prod.Close(); err != nil {
			panic(err)
		}
	}()

	return &AsyncProducer{prod, topic}
}

func (a AsyncProducer) Produce(ctx context.Context, message any) {
	b, err := json.Marshal(message)
	if err != nil {
		log.GetLogger(ctx).Error("failed to unmarshal", slog.Any("error", err.Error()))
	}

	k.AsyncProducer.Input() <- &sarama.ProducerMessage{Value: sarama.StringEncoder(b), Topic: k.topic}
}
