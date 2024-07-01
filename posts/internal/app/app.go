package app

import (
	"fmt"
	"github.com/MaksKazantsev/Chatter/posts/internal/config"
	"github.com/MaksKazantsev/Chatter/posts/internal/db/repository"
	userService "github.com/MaksKazantsev/Chatter/posts/internal/grpc"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/MaksKazantsev/Chatter/posts/internal/server"
	"github.com/MaksKazantsev/Chatter/posts/internal/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// New logger example
	l := log.InitLogger(cfg.Env)
	l.Info("Logger init success")

	// New db example
	repo := repository.NewRepository(repository.MustConnect(cfg.DB.GetAddr()))
	defer func() {
		_ = repo.Close()
	}()

	// Clients connection
	cl := userService.Connect(cfg.Services)

	// New service example
	srvc := service.NewService(repo)

	// New GRPC server
	srv := server.NewServer(l, srvc, cl)
	l.Info("All layers set up")

	shutdown(func() {
		l.Info("Server started on ", slog.Any("port: ", cfg.Port))
		run(srv, cfg.Port)
	})
	l.Info("Server shutting down...")
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
