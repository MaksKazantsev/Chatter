package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

type Auth interface {
	Register(ctx context.Context, req models.RegReq) (models.RegRes, error)
	Login(ctx context.Context, req models.LogReq) (string, string, error)
	Reset(ctx context.Context, req models.ResReq) error
}

func (s *service) Register(ctx context.Context, req models.RegReq) (models.RegRes, error) {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// hashing password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.RegRes{}, err
	}
	req.Password = hashed

	// generating id
	req.UUID = uuid.New().String()

	// generating tokens
	rToken, err := utils.NewToken(req.UUID, utils.REFRESH)
	if err != nil {
		return models.RegRes{}, err
	}
	aToken, err := utils.NewToken(req.UUID, utils.ACCESS)
	if err != nil {
		return models.RegRes{}, err
	}
	req.Refresh = rToken

	if err = s.repo.Register(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	code := strconv.Itoa(rand.Intn(9001) + 1000)

	go func() {
		if err = s.smtp.SendCode(code, req.Email); err != nil {
			fmt.Printf("smtp error: %v\n", err)
		}
	}()
	if err = s.repo.EmailAddCode(ctx, code, req.Email); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}
	return models.RegRes{
		UUID:         req.UUID,
		RefreshToken: rToken,
		AccessToken:  aToken,
	}, nil
}

func (s *service) Login(ctx context.Context, req models.LogReq) (string, string, error) {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// get password
	fmt.Println(req.Email)
	res, err := s.repo.GetHashAndID(ctx, req.Email)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	// compare passwords
	if err = utils.ComparePass(res.Password, req.Password); err != nil {
		return "", " ", fmt.Errorf("error: %w", err)
	}

	// generating token
	rToken, err := utils.NewToken(res.UUID, utils.REFRESH)
	if err != nil {
		return " ", " ", err
	}
	aToken, err := utils.NewToken(res.UUID, utils.ACCESS)
	if err != nil {
		return " ", " ", err
	}
	req.Refresh = rToken

	// login
	err = s.repo.Login(ctx, req)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	return rToken, aToken, nil
}
