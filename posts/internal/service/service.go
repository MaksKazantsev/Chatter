package service

import (
	"github.com/MaksKazantsev/Chatter/posts/internal/async"
	"github.com/MaksKazantsev/Chatter/posts/internal/db"
)

type Service struct {
	repo   db.Repository
	broker async.Producer
}

func NewService(repo db.Repository, broker async.Producer) *Service {
	return &Service{repo: repo, broker: broker}
}
