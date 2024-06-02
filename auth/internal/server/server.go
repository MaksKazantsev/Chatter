package server

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/service"
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pkg.UnimplementedUserServer
	log     log.Logger
	service service.Service
}

func NewServer(l log.Logger, service service.Service) *grpc.Server {
	srv := grpc.NewServer()
	pkg.RegisterUserServer(srv, &server{log: l, service: service})
	return srv
}

func (s *server) Register(ctx context.Context, req *pkg.RegisterReq) (*pkg.RegisterRes, error) {
	panic("implement me")
}

func (s *server) Login(ctx context.Context, req *pkg.LoginReq) (*pkg.LoginRes, error) {
	panic("implement me")
}

func (s *server) Reset(ctx context.Context, req *pkg.ResetReq) (*emptypb.Empty, error) {
	panic("implement me")
}
