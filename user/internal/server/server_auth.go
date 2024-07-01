package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) Register(ctx context.Context, req *pkg.RegisterReq) (*pkg.RegisterRes, error) {
	s.log.Debug("User microservice successfully received request")
	if err := s.validator.ValidateRegReq(req); err != nil {
		return nil, err
	}
	res, err := s.service.Auth.Register(log.WithLogger(ctx, s.log), s.converter.RegReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.RegResToPb(res), nil
}

func (s *server) Login(ctx context.Context, req *pkg.LoginReq) (*pkg.LoginRes, error) {
	s.log.Debug("User microservice successfully received request")
	if err := s.validator.ValidateLogReq(req); err != nil {
		return nil, err
	}
	rToken, aToken, err := s.service.Auth.Login(log.WithLogger(ctx, s.log), s.converter.LoginReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.LoginResToPb(aToken, rToken), nil
}

func (s *server) SendCode(ctx context.Context, req *pkg.SendReq) (*emptypb.Empty, error) {
	s.log.Debug("User microservice successfully received request")
	if err := s.validator.ValidateSendCodeReq(req); err != nil {
		return nil, err
	}
	if err := s.service.Auth.EmailSendCode(log.WithLogger(ctx, s.log), req.Email); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) VerifyCode(ctx context.Context, req *pkg.VerifyReq) (*pkg.VerifyRes, error) {
	s.log.Debug("User microservice successfully received request")
	if err := s.validator.ValidateVerifyCodeReq(req); err != nil {
		return nil, err
	}
	aToken, rToken, err := s.service.Auth.EmailVerifyCode(log.WithLogger(ctx, s.log), s.converter.VerifyCodeReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.VerifyCodeResToPb(aToken, rToken), nil
}

func (s *server) Recovery(ctx context.Context, req *pkg.RecoveryReq) (*emptypb.Empty, error) {
	s.log.Debug("User microservice successfully received request")
	if err := s.validator.ValidateRecoveryReq(req); err != nil {
		return nil, err
	}
	if err := s.service.Auth.PasswordRecovery(log.WithLogger(ctx, s.log), s.converter.RecoveryReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) UpdateToken(ctx context.Context, req *pkg.UpdateTokenReq) (*pkg.UpdateTokenRes, error) {
	s.log.Debug("User microservice successfully received request")
	aToken, rToken, err := s.service.Auth.UpdateTokens(log.WithLogger(ctx, s.log), req.RToken)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.UpdateTokensResToPb(aToken, rToken), nil
}

func (s *server) ParseToken(ctx context.Context, req *pkg.ParseTokenReq) (*pkg.ParseTokenRes, error) {
	s.log.Debug("ParseToken request received")
	id, err := s.service.Auth.ParseToken(log.WithLogger(ctx, s.log), req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return &pkg.ParseTokenRes{
		UUID: id,
	}, nil
}
