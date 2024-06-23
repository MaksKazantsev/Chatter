package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
)

type Friendship struct {
	repo db.Friendship
}

func NewFriendShip(repo db.Repository) *Friendship {
	return &Friendship{repo: repo}
}

func (f *Friendship) SuggestFriendship(ctx context.Context, req models.FriendShipReq) error {
	cl, err := utils.ParseToken(req.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	log.GetLogger(ctx).Debug("uc layer success ✔")
	if err = f.repo.SuggestFriendShip(ctx, cl["id"].(string), req.Receiver); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (f *Friendship) RefuseFriendship(ctx context.Context, req models.RefuseFriendShipReq) error {
	cl, err := utils.ParseToken(req.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	log.GetLogger(ctx).Debug("uc layer success ✔")
	if err = f.repo.RefuseFriendShip(ctx, cl["id"].(string), req.Sender); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (f *Friendship) AcceptFriendship(ctx context.Context, req models.AcceptFriendShipReq) error {
	_, err := utils.ParseToken(req.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	log.GetLogger(ctx).Debug("uc layer success ✔")

	return nil
}
