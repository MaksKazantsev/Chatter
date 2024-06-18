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
	Login(ctx context.Context, req models.LogReq) (string, string, error)
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
		return models.RegRes{}, err
	}
	req.Password = hashed

	// generating id
	req.UUID = uuid.New().String()

	if err = s.repo.Register(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	// generating tokens
	rToken, err := utils.NewToken(req.UUID, utils.REFRESH)
	if err != nil {
		return models.RegRes{}, err
	}
	aToken, err := utils.NewToken(req.UUID, utils.ACCESS)
	if err != nil {
		return models.RegRes{}, err
	}

	return models.RegRes{
		UUID:         req.UUID,
		RefreshToken: rToken,
		AccessToken:  aToken,
	}, nil
}

func (s *service) Login(ctx context.Context, req models.LogReq) (string, string, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// get password
	hash, id, err := s.repo.GetHashAndID(ctx, req.Email)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	// compare passwords
	if err = utils.ComparePass(hash, req.Password); err != nil {
		return "", " ", fmt.Errorf("error: %w", err)
	}

	// login
	err = s.repo.Login(ctx, req)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	// generating token
	rToken, err := utils.NewToken(id, utils.REFRESH)
	if err != nil {
		return " ", " ", err
	}
	aToken, err := utils.NewToken(id, utils.ACCESS)
	if err != nil {
		return " ", " ", err
	}
	return rToken, aToken, nil
}

func (s *service) Reset(ctx context.Context, req models.ResReq) error {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// parse token
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
	if err = utils.ComparePass(password, req.OldPassword); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// hashing new password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// reset password
	if err = s.repo.Reset(ctx, hashed, claims["id"].(string)); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
