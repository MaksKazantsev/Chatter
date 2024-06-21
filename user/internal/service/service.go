package service

import (
	"github.com/MaksKazantsev/Chatter/user/internal/db"
)

type Service struct {
	Auth       *Auth
	Friendship *Friendship
}

func NewService(repo db.Repository) *Service {
	return &Service{Auth: NewAuth(repo), Friendship: NewFriendShip(repo)}
}
