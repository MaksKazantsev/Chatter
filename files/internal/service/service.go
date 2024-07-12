package service

import (
	"github.com/MaksKazantsev/Chatter/files/internal/async"
	"github.com/MaksKazantsev/Chatter/files/internal/storage"
)

type Service struct {
	s3     storage.Storage
	broker async.Publisher
}

func NewService(strg storage.Storage, broker async.Publisher) *Service {
	return &Service{
		s3:     strg,
		broker: broker,
	}
}
