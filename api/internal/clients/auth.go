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
	SendCode(ctx context.Context, email string) error
	VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error)
	PasswordRecovery(ctx context.Context, req models.RecoveryReq) error
}

type userAuthCl struct {
	cl pkg.UserClient
	c  utils.Converter
}

func (u userAuthCl) PasswordRecovery(ctx context.Context, req models.RecoveryReq) error {
	_, err := u.cl.Recovery(ctx, u.c.RecoveryReqToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userAuthCl) VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error) {
	res, err := u.cl.VerifyCode(ctx, u.c.VerifyCodeReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u userAuthCl) SendCode(ctx context.Context, email string) error {
	_, err := u.cl.SendCode(ctx, u.c.SendCodeReqToPb(email))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
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

func NewUserAuth(cl pkg.UserClient) UserAuth {
	return &userAuthCl{
		cl: cl,
		c:  utils.NewConverter(),
	}
}
