package service

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

type Service interface {
	Register(ctx context.Context, req models.RegReq) (models.RegRes, error)
	Login(ctx context.Context, req models.LogReq) (string, string, error)
	UpdateTokens(ctx context.Context, refresh string) (string, string, error)
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
	EmailSendCode(ctx context.Context, email string) error
	EmailVerifyCode(ctx context.Context, code, email, t string) (string, string, error)
}

type service struct {
	repo db.Repository
	smtp utils.Smtp
}

func NewService(repo db.Repository) Service {
	return &service{repo: repo, smtp: utils.NewSmtp()}
}
