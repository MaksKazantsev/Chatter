package service

import (
	"github.com/MaksKazantsev/Chatter/files/internal/db"
	"github.com/MaksKazantsev/Chatter/files/internal/storage"
)

type Service struct {
	repo db.Repository
	s3   storage.Storage
}

func NewService(repo db.Repository, strg storage.Storage) *Service {
	return &Service{
		repo: repo,
		s3:   strg,
	}
}
