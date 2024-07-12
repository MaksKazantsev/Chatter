package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/async"
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
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

	b, err := json.Marshal(pkg.UpdateOnlineMessage{ID: msg.SenderID, LastOnline: time.Now()})
	if err != nil {
		return utils.NewError("failed to marshal message to broker "+err.Error(), utils.ErrInternal)
	}
	m.broker.Produce(ctx, pkg.Message{Type: "UpdateOnline", Data: b})

	return nil
}
func (m *Messages) DeleteMessage(ctx context.Context, messageID, id string) error {
	log.GetLogger(ctx).Info("Service layer success")

	if err := m.repo.DeleteMessage(ctx, messageID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	b, err := json.Marshal(pkg.UpdateOnlineMessage{ID: id, LastOnline: time.Now()})
	if err != nil {
		return utils.NewError("failed to marshal message to broker "+err.Error(), utils.ErrInternal)
	}
	m.broker.Produce(ctx, pkg.Message{Type: "UpdateOnline", Data: b})
	return nil
}

func (m *Messages) GetHistory(ctx context.Context, req models.GetHistoryReq, id string) ([]models.Message, error) {
	log.GetLogger(ctx).Info("Service layer success")

	b, err := json.Marshal(pkg.UpdateOnlineMessage{ID: id, LastOnline: time.Now()})
	if err != nil {
		return nil, utils.NewError("failed to marshal message to broker "+err.Error(), utils.ErrInternal)
	}
	m.broker.Produce(ctx, pkg.Message{Type: "UpdateOnline", Data: b})

	s := strings.Split(req.ChatID+id, "")
	sort.Strings(s)
	req.ChatID = strings.Join(s, "")

	res, err := m.repo.GetHistory(ctx, req, id)

	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}
	return res, nil
}
