package main

import (
	"github.com/MaksKazantsev/Chatter/api/internal/app"
	"github.com/MaksKazantsev/Chatter/api/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
