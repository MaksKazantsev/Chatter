package async

import (
	"context"
)

type Consumer interface {
	Consume(ctx context.Context) <-chan []byte
}

type ConsumerType int

const Kafka ConsumerType = iota

func NewConsumer(t ConsumerType, targetAddr, topic string) Consumer {
	switch t {
	case Kafka:
		return newKafka(targetAddr, topic)
	}
	return nil
}
