package main

import (
	"github.com/MaksKazantsev/Chatter/posts/internal/app"
	"github.com/MaksKazantsev/Chatter/posts/internal/config"
)

func main() {
	app.MustStart(config.MustInit())
}
