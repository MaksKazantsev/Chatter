package main

import (
	_ "github.com/MaksKazantsev/Chatter/api/docs"
	"github.com/MaksKazantsev/Chatter/api/internal/app"
	"github.com/MaksKazantsev/Chatter/api/internal/config"
)

// @title Social server API
// @version 1.0

func main() {
	app.MustStart(config.MustInit())
}
