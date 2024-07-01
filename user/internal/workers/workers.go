package workers

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

const (
	UpdateOnlineEvent = "updateonline"
)

type Message struct {
	Data []byte
}

type Worker interface {
	Start(ctx context.Context)
}

type worker struct {
	consumer sarama.Consumer
}

func NewWorker() Worker {
	_ = godotenv.Load(".env")
	c, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_ADDR")}, sarama.NewConfig())
	if err != nil {
		panic("Failed to init consumer: " + err.Error())
	}

	return &worker{consumer: c}
}

func (w *worker) Start(ctx context.Context) {
	partitions, _ := w.consumer.Partitions(UpdateOnlineEvent)
	c, err := w.consumer.ConsumePartition(UpdateOnlineEvent, partitions[0], sarama.OffsetNewest)
	if err != nil {
		panic("Failed to consume: " + err.Error())
	}

	messagesCh := c.Messages()
	logger := log.GetLogger(ctx)

	for i := 0; i < 5; i++ {
		go func() {
			for msg := range messagesCh {
				var m Message
				if err = json.Unmarshal(msg.Value, &m); err != nil {
					logger.Error("failed to unmarshal", slog.Any("error: ", err.Error()))
				}
				// TODO: finish
			}
		}()
	}
	select {
	case <-ctx.Done():
		_ = w.consumer.Close()
	}
}
