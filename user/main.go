package main

import (
	"github.com/MaksKazantsev/Chatter/user/internal/app"
	"github.com/MaksKazantsev/Chatter/user/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
