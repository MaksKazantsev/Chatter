package app

import (
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/config"
	"github.com/MaksKazantsev/SSO/auth/internal/db/repository"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/server"
	"github.com/MaksKazantsev/SSO/auth/internal/service"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// New logger example
	l := log.InitLogger(cfg.Env)
	l.Info("Logger init success")

	// New repository example
	repo := repository.NewRepository(repository.MustConnect(cfg.DB.GetAddr()))
	defer func() {
		_ = repo.Close()
	}()

	// New service example
	srvc := service.NewService(repo)

	// New GRPC server
	srv := server.NewServer(l, srvc)

	shutdown(func() {
		l.Info("server started")
		run(srv, cfg.Port)
	})
	l.Info("server stopped")
}

func run(srv *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic("failed to listen to tcp")
	}
	if err := srv.Serve(listener); err != nil {
		panic("failed to serve: " + err.Error())
	}
}

func shutdown(fn func()) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	go fn()
	<-chStop
}
