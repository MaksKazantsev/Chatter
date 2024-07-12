package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"

	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type UserClient interface {
	Register(ctx context.Context, req models.SignupReq) (string, string, error)
	Login(ctx context.Context, req models.LoginReq) (string, string, error)
	SendCode(ctx context.Context, email string) error
	VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error)
	PasswordRecovery(ctx context.Context, req models.RecoveryReq) error
	UpdateTokens(ctx context.Context, refresh string) (string, string, error)

	GetFriends(ctx context.Context, token string) ([]models.Friend, error)
	GetFs(ctx context.Context, token string) ([]models.FsReq, error)
	DeleteFriend(ctx context.Context, token, targetID string) error
	SuggestFs(ctx context.Context, token string, targetID string) error
	RefuseFs(ctx context.Context, token string, targetID string) error
	AcceptFs(ctx context.Context, token string, targetID string) error

	EditProfile(ctx context.Context, req models.UserProfileReq) error
	GetProfile(ctx context.Context, token, targetID string) (models.UserProfile, error)
	EditAvatar(ctx context.Context, token, avatar string) error
	DeleteAvatar(ctx context.Context, token string) error
}

type userClient struct {
	cl pkg.UserClient
	c  utils.Converter
}

func (u *userClient) UpdateTokens(ctx context.Context, refresh string) (string, string, error) {
	res, err := u.cl.UpdateToken(ctx, u.c.UpdateTokensToPb(refresh))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AToken, res.RToken, nil
}

func (u *userClient) PasswordRecovery(ctx context.Context, req models.RecoveryReq) error {
	_, err := u.cl.Recovery(ctx, u.c.RecoveryReqToPb(req))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u *userClient) VerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, string, error) {
	res, err := u.cl.VerifyCode(ctx, u.c.VerifyCodeReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u *userClient) SendCode(ctx context.Context, email string) error {
	_, err := u.cl.SendCode(ctx, u.c.SendCodeReqToPb(email))
	if err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (u *userClient) Register(ctx context.Context, req models.SignupReq) (string, string, error) {
	res, err := u.cl.Register(ctx, u.c.RegReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func (u *userClient) Login(ctx context.Context, req models.LoginReq) (string, string, error) {
	res, err := u.cl.Login(ctx, u.c.LogReqToPb(req))
	if err != nil {
		return "", "", utils.GRPCErrorToError(err)
	}
	return res.AccessToken, res.RefreshToken, nil
}

func NewUserAuth(cl pkg.UserClient) UserClient {
	return &userClient{
		cl: cl,
		c:  utils.NewConverter(),
	}
}
