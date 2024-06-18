package service

import (
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

type Service interface {
	Auth
	Internal
}

type service struct {
	repo db.Repository
	smtp utils.Smtp
}

func NewService(repo db.Repository) Service {
	return &service{repo: repo, smtp: utils.NewSmtp()}
}
