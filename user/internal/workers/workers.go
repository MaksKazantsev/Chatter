package workers

import (
	"context"
	"encoding/json"
	"github.com/MaksKazantsev/Chatter/user/internal/async"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
	"github.com/MaksKazantsev/Chatter/user/pkg"
)

const (
	UpdateOnlineType = "UpdateOnline"
	UpdateAvatarType = "UpdateAvatar"
)

type Worker interface {
	Start(ctx context.Context, maxWorkers int)
}

type worker struct {
	consumer       async.Consumer
	rabbitConsumer async.Consumer
	uc             *service.User
}

func NewWorker(cons async.Consumer, rabbitConsumer async.Consumer, uc *service.User) Worker {
	return &worker{consumer: cons, uc: uc, rabbitConsumer: rabbitConsumer}
}

func (w *worker) Start(ctx context.Context, maxWorkers int) {
	l := log.GetLogger(ctx)

	for i := 0; i < maxWorkers; i++ {
		go func() {
			var mes pkg.Message
			for msg := range w.consumer.Consume(ctx) {
				if err := json.Unmarshal(msg, &mes); err != nil {
					l.Error("failed to unmarshal: ", err.Error())
				}
				switch mes.Type {
				case UpdateOnlineType:
					var message models.UpdateOnlineMessage
					if err := json.Unmarshal(mes.Data, &message); err != nil {
						l.Error("failed to unmarshal: ", err.Error())
					}
					if err := w.updateOnline(ctx, message); err != nil {
						l.Error("failed to push message: ", err.Error())
					}
				}
			}
		}()
		go func() {
			var mes pkg.Message
			for msg := range w.rabbitConsumer.Consume(ctx) {
				if err := json.Unmarshal(msg, &mes); err != nil {
					l.Error("failed to unmarshal: ", err.Error())
				}
				switch mes.Type {
				case UpdateAvatarType:
					var message models.UpdateAvatarMessage
					if err := json.Unmarshal(mes.Data, &message); err != nil {
						l.Error("failed to unmarshal: ", err.Error())
					}
					if err := w.updateAvatar(ctx, message); err != nil {
						l.Error("failed to push message: ", err.Error())
					}
				}
			}
		}()
	}
}

func (w *worker) updateAvatar(ctx context.Context, mes models.UpdateAvatarMessage) error {
	l := log.GetLogger(ctx)

	if err := w.uc.EditAvatar(ctx, mes.ID, mes.Avatar); err != nil {
		l.Error("failed to update: ", err.Error())
	}
	return nil
}

func (w *worker) updateOnline(ctx context.Context, mes models.UpdateOnlineMessage) error {
	l := log.GetLogger(ctx)

	if err := w.uc.UpdateOnline(ctx, mes); err != nil {
		l.Error("failed to update: ", err.Error())
	}
	return nil
}
