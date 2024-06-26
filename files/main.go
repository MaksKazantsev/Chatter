package main

import (
	"github.com/MaksKazantsev/Chatter/files/internal/app"
	"github.com/MaksKazantsev/Chatter/files/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
