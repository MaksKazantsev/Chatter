package app

import (
	"github.com/MaksKazantsev/SSO/auth/internal/config"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
)

func MustStart(cfg *config.Config) {
	// New logger example
	l := log.InitLogger(cfg.Env)
	l.Info("Logger init success")
}

func run()
func shutdown()
