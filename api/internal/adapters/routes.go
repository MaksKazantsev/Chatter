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
	auth.Put("/login", ctrl.Auth.Login)
	auth.Put("/recovery", ctrl.Auth.PasswordRecovery)
	auth.Get("/email/verify", ctrl.Auth.VerifyCode)
	auth.Get("/email/send", ctrl.Auth.SendCode)
	auth.Get("/refresh", ctrl.Auth.UpdateTokens)
}
