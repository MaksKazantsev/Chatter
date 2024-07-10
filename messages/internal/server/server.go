package server

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/cache"
	userService "github.com/MaksKazantsev/Chatter/messages/internal/grpc"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/service"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pkg.UnimplementedMessagesServer
	log       log.Logger
	service   *service.Service
	converter utils.Converter
	validator utils.Validator
	userCl    userService.Clients
	cache     cache.Cache
}

func NewServer(l log.Logger, service *service.Service, userCl userService.Clients, cache cache.Cache) *grpc.Server {
	srv := grpc.NewServer()
	pkg.RegisterMessagesServer(srv, &server{log: l, service: service, validator: utils.NewValidator(), converter: utils.NewConverter(), userCl: userCl, cache: cache})
	return srv
}
