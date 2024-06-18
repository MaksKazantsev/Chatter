package app

import (
	"fmt"
	"github.com/MaksKazantsev/SSO/api/internal/adapters"
	"github.com/MaksKazantsev/SSO/api/internal/clients"
	"github.com/MaksKazantsev/SSO/api/internal/config"
	"github.com/MaksKazantsev/SSO/api/internal/log"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// New logger example
	l := log.InitLogger(cfg.Env)
	l.Info("Logger init success")

	cl := clients.Connect(cfg.Services)

	// New controller
	ctrl := adapters.NewController(cl)

	// New app
	app := fiber.New()

	// Init routes
	adapters.InitRoutes(app, ctrl)

	shutdown(func() {
		l.Info("server started")
		run(app, cfg.Port)
	})
	l.Info("server stopped")
}

func run(srv *fiber.App, port string) {
	if err := srv.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic("failed to listen: " + err.Error())
	}
}

func shutdown(fn func()) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	go fn()
	<-chStop
}
