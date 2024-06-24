package service

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
)

type Service struct {
	Messages *Messages
}

func NewService(repo db.Repository) *Service {
	return &Service{Messages: NewMessages(repo)}
}
