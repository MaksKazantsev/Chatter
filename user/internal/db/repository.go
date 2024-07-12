package db

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
)

type Repository interface {
	User
	Auth
}

type Auth interface {
	Login(ctx context.Context, req models.LogReq) error
	Register(ctx context.Context, req models.RegReq) error
	EmailAddCode(ctx context.Context, code string, email string) error
	EmailVerifyCode(ctx context.Context, req models.VerifyCodeReq) (string, error)
	PasswordRecovery(ctx context.Context, cr models.Credentials) error
	UpdateRToken(ctx context.Context, id, rToken string) error
	GetHashAndID(ctx context.Context, email string) (HashAndID, error)
}
type User interface {
	UpdateOnline(ctx context.Context, mes models.UpdateOnlineMessage) error
	EditProfile(ctx context.Context, profile models.UserProfile) error
	SuggestFs(ctx context.Context, reqID, receiverID, senderID string) error
	RefuseFs(ctx context.Context, reqID, receiverID string) error
	GetFs(ctx context.Context, receiverID string) ([]models.FsReq, error)
	AcceptFs(ctx context.Context, senderID, receiverID, reqID string) error
	DeleteFriend(ctx context.Context, userID, friendID string) error
	GetFriends(ctx context.Context, userID string) ([]models.Friend, error)
	GetProfile(ctx context.Context, userID string) (models.GetUserProfile, error)
	EditAvatar(ctx context.Context, userID, avatar string) error
	DeleteAvatar(ctx context.Context, userID string) error
}

type HashAndID struct {
	Password string `db:"password" json:"password"`
	UUID     string `db:"uuid" json:"uuid"`
}
