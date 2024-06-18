package db

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
)

type Repository interface {
	Auth
	Internal
	Verification
}

type Auth interface {
	Login(ctx context.Context, req models.LogReq) error
	Register(ctx context.Context, req models.RegReq) error
}
type Internal interface {
	GetHashAndID(ctx context.Context, email string) (HashAndID, error)
	EmailAddCode(ctx context.Context, code, email string) error
	EmailVerifyCode(ctx context.Context, code, email, t string) error
}

type Verification interface {
	EmailAddCode(ctx context.Context, code, email string) error
	EmailVerifyCode(ctx context.Context, code, email, t string) error
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
}

type HashAndID struct {
	Password string `db:"password" json:"password"`
	UUID     string `db:"uuid" json:"uuid"`
}
