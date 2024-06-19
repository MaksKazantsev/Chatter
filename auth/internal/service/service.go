package service

import (
	"github.com/MaksKazantsev/SSO/auth/internal/db"
)

type Service struct {
	Auth       Auth
	FriendShip FriendShip
}

func NewService(repo db.Repository) *Service {
	return &Service{Auth: NewAuth(repo), FriendShip: NewFriendShip(repo)}
}
