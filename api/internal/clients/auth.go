package clients

import (
	"context"
	"github.com/MaksKazantsev/SSO/api/internal/models"
	"github.com/MaksKazantsev/SSO/api/internal/utils"
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
)

type UserAuth interface {
	Register(ctx context.Context, req models.SignupReq) (string, string, error)
	Login(ctx context.Context, req models.LoginReq) (string, string, error)
	Reset(ctx context.Context, req models.ResetReq) error
}

type userAuthCl struct {
	cl pkg.UserClient
	c  utils.Converter
}

func (u userAuthCl) Register(ctx context.Context, req models.SignupReq) (string, string, error) {
	res, err := u.cl.Register(ctx, u.c.RegReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u userAuthCl) Login(ctx context.Context, req models.LoginReq) (string, string, error) {
	res, err := u.cl.Login(ctx, u.c.LogReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u userAuthCl) Reset(ctx context.Context, req models.ResetReq) error {
	_, err := u.cl.Reset(ctx, u.c.ResReqToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func NewUserAuth(cl pkg.UserClient) UserAuth {
	return &userAuthCl{
		cl: cl,
		c:  utils.NewConverter(),
	}
}
