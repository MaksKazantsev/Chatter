package server

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	"github.com/MaksKazantsev/Chatter/user/internal/utils/converter"
	"github.com/MaksKazantsev/Chatter/user/internal/utils/validator"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pkg.UnimplementedUserServer
	log       log.Logger
	service   service.Service
	converter converter.Converter
	validator validator.Validator
}

func NewServer(l log.Logger, service service.Service) *grpc.Server {
	srv := grpc.NewServer()
	pkg.RegisterUserServer(srv, &server{log: l, service: service, validator: validator.NewValidator(), converter: converter.NewConverter()})
	return srv
}

func (s *server) Register(ctx context.Context, req *pkg.RegisterReq) (*pkg.RegisterRes, error) {
	if err := s.validator.ValidateRegReq(req); err != nil {
		return nil, err
	}
	res, err := s.service.Register(log.WithLogger(ctx, s.log), s.converter.RegReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.RegResToPb(res), nil
}

func (s *server) Login(ctx context.Context, req *pkg.LoginReq) (*pkg.LoginRes, error) {
	if err := s.validator.ValidateLogReq(req); err != nil {
		return nil, err
	}
	rToken, aToken, err := s.service.Login(log.WithLogger(ctx, s.log), s.converter.LoginReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.LoginResToPb(aToken, rToken), nil
}

func (s *server) SendCode(ctx context.Context, req *pkg.SendReq) (*emptypb.Empty, error) {
	if err := s.validator.ValidateSendCodeReq(req); err != nil {
		return nil, err
	}
	if err := s.service.EmailSendCode(log.WithLogger(ctx, s.log), s.converter.SendCodeReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) VerifyCode(ctx context.Context, req *pkg.VerifyReq) (*pkg.VerifyRes, error) {
	if err := s.validator.ValidateVerifyCodeReq(req); err != nil {
		return nil, err
	}
	code, email, t := s.converter.VerifyCodeReqToService(req)
	aToken, rToken, err := s.service.EmailVerifyCode(log.WithLogger(ctx, s.log), code, email, t)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.VerifyCodeResToPb(aToken, rToken), nil
}

func (s *server) Recovery(ctx context.Context, req *pkg.RecoveryReq) (*emptypb.Empty, error) {
	if err := s.validator.ValidateRecoveryReq(req); err != nil {
		return nil, err
	}
	if err := s.service.PasswordRecovery(log.WithLogger(ctx, s.log), s.converter.RecoveryReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) UpdateToken(ctx context.Context, req *pkg.UpdateTokenReq) (*pkg.UpdateTokenRes, error) {
	aToken, rToken, err := s.service.UpdateTokens(log.WithLogger(ctx, s.log), s.converter.UpdateTokensReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.UpdateTokensResToPb(aToken, rToken), nil
}
