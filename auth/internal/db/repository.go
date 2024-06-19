package db

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
)

type Repository interface {
	Auth
	Friendship
}

type Auth interface {
	Login(ctx context.Context, req models.LogReq) error
	Register(ctx context.Context, req models.RegReq) error
	EmailAddCode(ctx context.Context, code string, email string) error
	GetHashAndID(ctx context.Context, email string) (HashAndID, error)
	EmailVerifyCode(ctx context.Context, code, email, t string) (string, error)
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
	UpdateRToken(ctx context.Context, id, rToken string) error
}
type Friendship interface {
}

type HashAndID struct {
	Password string `db:"password" json:"password"`
	UUID     string `db:"uuid" json:"uuid"`
}
