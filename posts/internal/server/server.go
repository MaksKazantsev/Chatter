package server

import (
	"github.com/MaksKazantsev/Chatter/posts/internal/cache"
	userService "github.com/MaksKazantsev/Chatter/posts/internal/grpc"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/MaksKazantsev/Chatter/posts/internal/service"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils/converter"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils/validator"
	pkg "github.com/MaksKazantsev/Chatter/posts/pkg/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pkg.UnimplementedPostsServer
	log    log.Logger
	srvc   *service.Service
	c      converter.Converter
	v      validator.Validator
	userCl userService.Clients
	cache  cache.Cache
}

func NewServer(l log.Logger, srvc *service.Service, userCl userService.Clients, cache cache.Cache) *grpc.Server {
	srv := grpc.NewServer()
	pkg.RegisterPostsServer(srv, &server{log: l, srvc: srvc, userCl: userCl, c: converter.NewConverter(), v: validator.NewValidator(), cache: cache})
	return srv
}
