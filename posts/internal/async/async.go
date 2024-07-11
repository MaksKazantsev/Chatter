package async

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/segmentio/kafka-go"
	"time"
)

var _ Producer = &producer{}

type producer struct {
	*kafka.Writer
	topic string
}

func NewProducer(addr string, topic string) Producer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
		Async:        true,
	}

	return &producer{writer, topic}
}

func (a *producer) Produce(ctx context.Context, message any) {
	b, err := json.Marshal(message)
	if err != nil {
		log.GetLogger(ctx).Error("producing error: ", err.Error())
	}
	fmt.Println("p")
	err = a.WriteMessages(ctx, kafka.Message{Value: b})
}
