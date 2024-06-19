package server

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/service"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
	"github.com/MaksKazantsev/SSO/auth/internal/utils/converter"
	"github.com/MaksKazantsev/SSO/auth/internal/utils/validator"
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pkg.UnimplementedUserServer
	log       log.Logger
	service   *service.Service
	converter converter.Converter
	validator validator.Validator
}

func NewServer(l log.Logger, service *service.Service) *grpc.Server {
	srv := grpc.NewServer()
	pkg.RegisterUserServer(srv, &server{log: l, service: service, validator: validator.NewValidator(), converter: converter.NewConverter()})
	return srv
}

func (s *server) Register(ctx context.Context, req *pkg.RegisterReq) (*pkg.RegisterRes, error) {
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
	if err := s.validator.ValidateSendCodeReq(req); err != nil {
		return nil, err
	}
	if err := s.service.Auth.EmailSendCode(log.WithLogger(ctx, s.log), s.converter.SendCodeReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *server) VerifyCode(ctx context.Context, req *pkg.VerifyReq) (*pkg.VerifyRes, error) {
	if err := s.validator.ValidateVerifyCodeReq(req); err != nil {
		return nil, err
	}
	code, email, t := s.converter.VerifyCodeReqToService(req)
	aToken, rToken, err := s.service.Verification.EmailVerifyCode(log.WithLogger(ctx, s.log), code, email, t)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return s.converter.VerifyCodeResToPb(aToken, rToken), nil
}

func (s *server) Recovery(ctx context.Context, req *pkg.RecoveryReq) (*emptypb.Empty, error) {
	if err := s.validator.ValidateRecoveryReq(req); err != nil {
		return nil, err
	}
	if err := s.service.Verification.PasswordRecovery(log.WithLogger(ctx, s.log), s.converter.RecoveryReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}
	return &emptypb.Empty{}, nil
}
