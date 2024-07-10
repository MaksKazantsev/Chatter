package app

import (
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/async"
	"github.com/MaksKazantsev/Chatter/messages/internal/cache"
	"github.com/MaksKazantsev/Chatter/messages/internal/config"
	"github.com/MaksKazantsev/Chatter/messages/internal/db/repository"
	userService "github.com/MaksKazantsev/Chatter/messages/internal/grpc"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/server"
	"github.com/MaksKazantsev/Chatter/messages/internal/service"
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

	// New db example
	repo := repository.NewRepository(repository.MustConnect(cfg.DB.GetAddr()))
	defer func() {
		_ = repo.Close()
	}()

	// External communication
	prod := async.NewProducer(os.Getenv("KAFKA_ADDR"), os.Getenv("KAFKA_TOPIC"))
	cl := userService.Connect(cfg.Services)

	// New service example
	srvc := service.NewService(repo, prod)

	// New Cache example
	c := cache.NewRedis()

	// New GRPC server
	srv := server.NewServer(l, srvc, cl, c)
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
