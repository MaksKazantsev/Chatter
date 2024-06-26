package server

import (
	userService "github.com/MaksKazantsev/Chatter/files/internal/grpc"
	"github.com/MaksKazantsev/Chatter/files/internal/log"
	"github.com/MaksKazantsev/Chatter/files/internal/service"
	"github.com/MaksKazantsev/Chatter/files/internal/utils/converter"
	"github.com/MaksKazantsev/Chatter/files/internal/utils/validator"
	filesPkg "github.com/MaksKazantsev/Chatter/files/pkg/grpc"
	"google.golang.org/grpc"
)

type server struct {
	filesPkg.UnimplementedFilesServer
	service *service.Service
	userCl  userService.Clients
	l       log.Logger
	v       validator.Validator
	c       converter.Converter
}

func NewServer(service *service.Service, l log.Logger, userCl userService.Clients) *grpc.Server {
	srv := grpc.NewServer()
	filesPkg.RegisterFilesServer(srv, &server{l: l, service: service, v: validator.NewValidator(), c: converter.NewConverter(), userCl: userCl})
	return srv
}
