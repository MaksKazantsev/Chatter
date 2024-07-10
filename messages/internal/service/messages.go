package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/async"
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	"github.com/MaksKazantsev/Chatter/messages/pkg"
	"github.com/google/uuid"
	"sort"
	"strings"
	"time"
)

type Messages struct {
	repo   db.Messages
	broker async.Producer
}

func NewMessages(repo db.Messages, broker async.Producer) *Messages {
	return &Messages{repo: repo, broker: broker}
}

func (m *Messages) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {
	msg.MessageID = uuid.New().String()
	msg.SentAt = time.Now()

	log.GetLogger(ctx).Info("Service layer success")

	if err := m.repo.CreateMessage(ctx, msg, receiverOffline); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	m.broker.Produce(ctx, pkg.KafkaMessage{ID: msg.SenderID, LastOnline: time.Now()})

	return nil
}
func (m *Messages) DeleteMessage(ctx context.Context, messageID, id string) error {
	log.GetLogger(ctx).Info("Service layer success")

	if err := m.repo.DeleteMessage(ctx, messageID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	m.broker.Produce(ctx, pkg.KafkaMessage{ID: id, LastOnline: time.Now()})
	return nil
}

func (m *Messages) GetHistory(ctx context.Context, req models.GetHistoryReq, uuid string) ([]models.Message, error) {
	log.GetLogger(ctx).Info("Service layer success")

	m.broker.Produce(ctx, pkg.KafkaMessage{ID: uuid, LastOnline: time.Now()})

	s := strings.Split(req.ChatID+uuid, "")
	sort.Strings(s)
	req.ChatID = strings.Join(s, "")

	res, err := m.repo.GetHistory(ctx, req, uuid)

	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}
	return res, nil
}
