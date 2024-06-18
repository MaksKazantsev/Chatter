package main

import (
	"github.com/MaksKazantsev/SSO/api/internal/app"
	"github.com/MaksKazantsev/SSO/api/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
