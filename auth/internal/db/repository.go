package db

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
)

type Repository interface {
	Login(ctx context.Context, req models.LogReq) error
	Register(ctx context.Context, req models.RegReq) error

	EmailVerifyCode(ctx context.Context, code, email, t string) (string, error)
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
	EmailAddCode(ctx context.Context, code string, email string) error

	GetHashAndID(ctx context.Context, email string) (HashAndID, error)
	UpdateRToken(ctx context.Context, id, rToken string) error
}

type HashAndID struct {
	Password string `db:"password" json:"password"`
	UUID     string `db:"uuid" json:"uuid"`
}
