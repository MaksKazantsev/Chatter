package async

import (
	"context"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
	"os"
)

type kafka struct {
	consumer *kafkago.Reader
	topic    string
}

func newKafka(targetAddr, topic string) *kafka {
	cfg := kafkago.ReaderConfig{
		Brokers: []string{targetAddr},
		Topic:   topic,
	}

	cons := kafkago.NewReader(cfg)

	return &kafka{
		consumer: cons,
		topic:    topic,
	}
}

func (k *kafka) Consume(ctx context.Context) <-chan []byte {
	ch := make(chan []byte, 1)
	go func() {
		defer close(ch)
		defer func() { _ = k.consumer.Close() }()

		for {
			msg, err := k.consumer.ReadMessage(ctx)
			fmt.Println(msg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error consuming: %v\n", err)
				break
			}

			ch <- msg.Value
		}
	}()

	return ch
}
