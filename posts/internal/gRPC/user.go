package gRPC

import (
	"context"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type User interface {
	ParseToken(ctx context.Context, token string) (string, error)
}

type userCl struct {
	cl pkg.UserClient
}

func NewUser(cl pkg.UserClient) User {
	return &userCl{cl: cl}
}

func (u userCl) ParseToken(ctx context.Context, token string) (string, error) {
	res, err := u.cl.ParseToken(ctx, &pkg.ParseTokenReq{Token: token})
	if err != nil {
		return "", err
	}
	return res.UUID, nil
}
