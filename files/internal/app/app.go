package app

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/files/internal/async"
	"github.com/MaksKazantsev/Chatter/files/internal/config"
	userService "github.com/MaksKazantsev/Chatter/files/internal/grpc"
	"github.com/MaksKazantsev/Chatter/files/internal/log"
	"github.com/MaksKazantsev/Chatter/files/internal/server"
	"github.com/MaksKazantsev/Chatter/files/internal/service"
	"github.com/MaksKazantsev/Chatter/files/internal/storage/s3"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// Load .env
	if err := godotenv.Load(); err != nil {
		panic("failed to load env: " + err.Error())
	}

	// New logger example
	l := log.InitLogger(cfg.Env)
	l.Info("Logger init success")

	// External communications
	cl := userService.Connect(cfg.Services)
	pub := async.NewPublisher(log.WithLogger(context.Background(), l), os.Getenv("RABBITMQ_ADDR"))

	// New service example
	srvc := service.NewService(s3.NewStorage(), pub)

	// New GRPC server
	srv := server.NewServer(srvc, l, cl)
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
