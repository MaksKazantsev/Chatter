package db

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
)

type Repository interface {
	Auth
	Internal
}

type Auth interface {
	Login(ctx context.Context, req models.LogReq) (LoginInfo, error)
	Register(ctx context.Context, req models.RegReq) error
	Reset(ctx context.Context, password, uuid string) error
}
type Internal interface {
	GetPasswordByEmail(ctx context.Context, email string) (string, error)
	GetPasswordByUUID(ctx context.Context, uuid string) (string, error)
}

type LoginInfo struct {
	UUID string
}
