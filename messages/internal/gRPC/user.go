package gRPC

import (
	"context"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type User interface {
	ParseToken(ctx context.Context, token string) (string, error)
}

type userCl struct {
	cl pkg.UserClient
	c  utils.Converter
}

func NewUser(cl pkg.UserClient) User {
	return &userCl{
		cl: cl,
		c:  utils.NewConverter(),
	}
}

func (u userCl) ParseToken(ctx context.Context, token string) (string, error) {
	res, err := u.cl.ParseToken(ctx, u.c.ParseTokenToPb(token))
	if err != nil {
		return "", err
	}
	return res.UUID, nil
}
