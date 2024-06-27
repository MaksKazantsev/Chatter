package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	"sort"
	"strings"
)

type User struct {
	repo db.User
}

func NewUser(repo db.User) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) EditProfile(ctx context.Context, req models.UserProfile) error {
	log.GetLogger(ctx).Info("Service layer success")
	if err := u.repo.EditProfile(ctx, req); err != nil {
		return fmt.Errorf("repo error: %w, err")
	}
	return nil
}

func (u *User) SuggestFs(ctx context.Context, senderID, receiverID string) error {
	if senderID == receiverID {
		return utils.NewError("you cant send friend req to your own", utils.ErrBadRequest)
	}

	s := strings.Split(senderID+receiverID, "")
	sort.Strings(s)
	reqID := strings.Join(s, "")

	log.GetLogger(ctx).Info("Service layer success")
	if err := u.repo.SuggestFs(ctx, reqID, receiverID, senderID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) RefuseFs(ctx context.Context, senderID, receiverID string) error {
	s := strings.Split(senderID+receiverID, "")
	sort.Strings(s)
	reqID := strings.Join(s, "")

	log.GetLogger(ctx).Info("Service layer success")
	if err := u.repo.RefuseFs(ctx, reqID, receiverID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) GetFs(ctx context.Context, receiverID string) ([]models.FsReq, error) {
	log.GetLogger(ctx).Info("Service layer success")
	res, err := u.repo.GetFs(ctx, receiverID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return res, nil
}
