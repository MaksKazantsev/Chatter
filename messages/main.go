package main

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/app"
	"github.com/MaksKazantsev/Chatter/messages/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
