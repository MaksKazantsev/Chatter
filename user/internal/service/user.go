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
	"time"
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
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: req.UUID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
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

	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: senderID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	if err := u.repo.SuggestFs(ctx, reqID, receiverID, senderID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) RefuseFs(ctx context.Context, senderID, receiverID string) error {
	s := strings.Split(senderID+receiverID, "")
	sort.Strings(s)
	reqID := strings.Join(s, "")

	log.GetLogger(ctx).Debug("Service layer success")
	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: receiverID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	if err := u.repo.RefuseFs(ctx, reqID, receiverID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) GetFs(ctx context.Context, receiverID string) ([]models.FsReq, error) {
	log.GetLogger(ctx).Debug("Service layer success")
	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: receiverID, LastOnline: time.Now()}); err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}
	res, err := u.repo.GetFs(ctx, receiverID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return res, nil
}

func (u *User) AcceptFs(ctx context.Context, senderID, receiverID string) error {
	s := strings.Split(senderID+receiverID, "")
	sort.Strings(s)
	reqID := strings.Join(s, "")

	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: receiverID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	if err := u.repo.AcceptFs(ctx, senderID, receiverID, reqID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) DeleteFriend(ctx context.Context, userID, friendID string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: userID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	if err := u.repo.DeleteFriend(ctx, userID, friendID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (u *User) GetFriends(ctx context.Context, userID string) ([]models.Friend, error) {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: userID, LastOnline: time.Now()}); err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	res, err := u.repo.GetFriends(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}
	return res, nil
}

func (u *User) GetProfile(ctx context.Context, userID string) (models.GetUserProfile, error) {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: userID, LastOnline: time.Now()}); err != nil {
		return models.GetUserProfile{}, fmt.Errorf("repo error: %w", err)
	}

	profile, err := u.repo.GetProfile(ctx, userID)
	if err != nil {
		return models.GetUserProfile{}, fmt.Errorf("repo error: %w", err)
	}
	return profile, nil
}

func (u *User) EditAvatar(ctx context.Context, userID, avatar string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: userID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	err := u.repo.EditAvatar(ctx, userID, avatar)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (u *User) DeleteAvatar(ctx context.Context, userID string) error {
	log.GetLogger(ctx).Debug("Service layer success")

	if err := u.repo.UpdateOnline(ctx, models.UpdateOnlineMessage{ID: userID, LastOnline: time.Now()}); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	err := u.repo.DeleteAvatar(ctx, userID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (u *User) UpdateOnline(ctx context.Context, mes models.UpdateOnlineMessage) error {
	// logging
	log.GetLogger(ctx).Debug("Service layer success")
	if err := u.repo.UpdateOnline(ctx, mes); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
