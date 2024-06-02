package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
	"github.com/google/uuid"
)

type Service interface {
	Register(ctx context.Context, req models.RegReq) (models.RegRes, error)
	Login(ctx context.Context, req models.LogReq) (string, error)
	Reset(ctx context.Context, req models.ResReq) error
}

type service struct {
	repo db.Repository
}

func NewService(repo db.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(ctx context.Context, req models.RegReq) (models.RegRes, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// hashing password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.RegRes{}, fmt.Errorf("failed to hash password: %w", err)
	}
	req.Password = hashed

	// generating id
	req.UUID = uuid.New().String()

	if err = s.repo.Register(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	// generating token
	token, err := utils.NewToken(req.UUID)
	if err != nil {
		return models.RegRes{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return models.RegRes{
		UUID:  req.UUID,
		Token: token,
	}, nil
}

func (s *service) Login(ctx context.Context, req models.LogReq) (string, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// login
	info, err := s.repo.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}

	// compare passwords
	if err = utils.ComparePass(req.Password, info.Password); err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	// generating token
	token, err := utils.NewToken(info.UUID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}

func (s *service) Reset(ctx context.Context, req models.ResReq) error {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	claims, err := utils.ParseToken(req.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	// get password
	password, err := s.repo.GetPasswordByUUID(ctx, claims["id"].(string))
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	// compare password
	if err = utils.ComparePass(req.Password, password); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// hashing new password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err = s.repo.Reset(ctx, hashed, claims["id"].(string)); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
