package service

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/async"
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
)

type Service struct {
	Messages *Messages
}

func NewService(repo db.Repository, producer async.Producer) *Service {
	return &Service{Messages: NewMessages(repo, producer)}
}
