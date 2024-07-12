package app

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/async"
	"github.com/MaksKazantsev/Chatter/user/internal/config"
	"github.com/MaksKazantsev/Chatter/user/internal/db/repository"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/server"
	"github.com/MaksKazantsev/Chatter/user/internal/service"
	"github.com/MaksKazantsev/Chatter/user/internal/workers"
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
	l.Info("Database layer set up")

	// New service example
	srvc := service.NewService(repo)
	l.Info("Service layer set up")

	// Init task workers
	w := workers.NewWorker(async.NewConsumer(async.Kafka, os.Getenv("KAFKA_ADDR"), os.Getenv("KAFKA_TOPIC")), async.NewConsumer(async.Rabbit, os.Getenv("RABBITMQ_ADDR"), os.Getenv("RABBITMQ_ROUTING_KEY")), srvc.User)

	// New GRPC server
	srv := server.NewServer(l, srvc)
	l.Info("All layers set up")

	shutdown(func() {
		w.Start(log.WithLogger(context.Background(), l), cfg.Worker.MaxWorkers)
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
	if err = srv.Serve(listener); err != nil {
		panic("failed to serve: " + err.Error())
	}
}

// graceful shutdown
func shutdown(fn func()) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	go fn()
	<-chStop
}
