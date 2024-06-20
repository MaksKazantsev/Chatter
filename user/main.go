package main

import (
	"github.com/MaksKazantsev/SSO/auth/internal/app"
	"github.com/MaksKazantsev/SSO/auth/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
