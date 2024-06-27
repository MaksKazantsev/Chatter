package service

import (
	"github.com/MaksKazantsev/Chatter/user/internal/db"
)

type Service struct {
	Auth *Auth
	User *User
}

func NewService(repo db.Repository) *Service {
	return &Service{Auth: NewAuth(repo), User: NewUser(repo)}
}
