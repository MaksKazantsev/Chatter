package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/async"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
)

type Message struct {
	ID string `json:"id"`
}

type Worker interface {
	Start(ctx context.Context, maxWorkers int)
}

type worker struct {
	consumer async.Consumer
	uc       *service.Auth
}

func NewWorker(cons async.Consumer, uc *service.Auth) Worker {
	return &worker{consumer: cons, uc: uc}
}

func (w *worker) Start(ctx context.Context, maxWorkers int) {
	l := log.GetLogger(ctx)

	fmt.Println(1)

	for i := 0; i < maxWorkers; i++ {
		go func() {
			var id string
			for msg := range w.consumer.Consume(ctx) {
				if err := json.Unmarshal(msg, &id); err != nil {
					l.Error("failed to unmarshal: ", err.Error())
				}
				if err := w.updateOnline(ctx, id); err != nil {
					l.Error("failed to push message: ", err.Error())
				}
			}
		}()
	}
}

func (w *worker) updateOnline(ctx context.Context, id string) error {
	l := log.GetLogger(ctx)
	if err := w.uc.UpdateOnline(ctx, id); err != nil {
		l.Error("failed to update: ", err.Error())
	}
	return nil
}
