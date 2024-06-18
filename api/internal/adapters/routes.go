package adapters

import (
	"github.com/MaksKazantsev/SSO/api/internal/adapters/handlers"
	"github.com/MaksKazantsev/SSO/api/internal/clients"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Auth *handlers.Auth
}

func NewController(clients clients.Clients) *Controller {
	return &Controller{
		Auth: handlers.NewAuth(clients.UserClient),
	}
}

func InitRoutes(app *fiber.App, ctrl *Controller) {
	auth := app.Group("/auth")
	auth.Post("/register", ctrl.Auth.Register)
	auth.Get("/login", ctrl.Auth.Login)
	auth.Put("/reset", ctrl.Auth.Reset)
}
