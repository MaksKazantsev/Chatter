package service

import (
	"github.com/MaksKazantsev/SSO/auth/internal/db"
)

type Service struct {
	Auth         *Auth
	Verification *Verification
}

func NewService(repo db.Repository) *Service {
	return &Service{Auth: NewAuth(repo), Verification: NewVerification(repo)}
}
