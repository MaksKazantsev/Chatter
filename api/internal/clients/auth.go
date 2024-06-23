package clients

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"

	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type User interface {
	Register(ctx context.Context, req models.SignupReq) (string, string, error)
	Login(ctx context.Context, req models.LoginReq) (string, string, error)
	SendCode(ctx context.Context, email string) error
	VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error)
	PasswordRecovery(ctx context.Context, req models.RecoveryReq) error
	UpdateTokens(ctx context.Context, refresh string) (string, string, error)
	SuggestFriendShip(ctx context.Context, req models.FriendShipReq) error
	RefuseFriendShip(ctx context.Context, req models.RefuseFriendShipReq) error
	ParseToken(ctx context.Context, token string) (string, error)
}

type userCl struct {
	cl pkg.UserClient
	c  utils.Converter
}

func (u userCl) ParseToken(ctx context.Context, token string) (string, error) {
	res, err := u.cl.ParseToken(ctx, u.c.ParseTokenToPb(token))
	if err != nil {
		return "", utils.GRPCErrorToError(err)
	}
	return res.UUID, nil
}

func (u userCl) UpdateTokens(ctx context.Context, refresh string) (string, string, error) {
	res, err := u.cl.UpdateToken(ctx, u.c.UpdateTokensToPb(refresh))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AToken, res.RToken, nil
}

func (u userCl) PasswordRecovery(ctx context.Context, req models.RecoveryReq) error {
	_, err := u.cl.Recovery(ctx, u.c.RecoveryReqToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error) {
	res, err := u.cl.VerifyCode(ctx, u.c.VerifyCodeReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u userCl) SendCode(ctx context.Context, email string) error {
	_, err := u.cl.SendCode(ctx, u.c.SendCodeReqToPb(email))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u userCl) Register(ctx context.Context, req models.SignupReq) (string, string, error) {
	fmt.Println("1")
	res, err := u.cl.Register(ctx, u.c.RegReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u userCl) Login(ctx context.Context, req models.LoginReq) (string, string, error) {
	res, err := u.cl.Login(ctx, u.c.LogReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func NewUserAuth(cl pkg.UserClient) User {
	return &userCl{
		cl: cl,
		c:  utils.NewConverter(),
	}
}
