package server

import (
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
	"github.com/MaksKazantsev/Chatter/user/internal/utils/converter"
	"github.com/MaksKazantsev/Chatter/user/internal/utils/validator"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc"
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
