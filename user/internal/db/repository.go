package db

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
)

type Repository interface {
	Friendship
	Auth
}

type Auth interface {
	Login(ctx context.Context, req models.LogReq) error
	Register(ctx context.Context, req models.RegReq) error
	EmailAddCode(ctx context.Context, code string, email string) error
	EmailVerifyCode(ctx context.Context, code, email, t string) (string, error)
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
	UpdateRToken(ctx context.Context, id, rToken string) error
	GetHashAndID(ctx context.Context, email string) (HashAndID, error)
	UpdateOnline(ctx context.Context, uuid string) error
}
type Friendship interface {
}

type HashAndID struct {
	Password string `db:"password" json:"password"`
	UUID     string `db:"uuid" json:"uuid"`
}
