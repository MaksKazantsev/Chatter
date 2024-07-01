package service

import "github.com/MaksKazantsev/Chatter/posts/internal/db"

type Service struct {
	repo db.Repository
}

func NewService(repo db.Repository) *Service {
	return &Service{repo: repo}
}
