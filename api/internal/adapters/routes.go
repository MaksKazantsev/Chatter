package adapters

import (
	"github.com/MaksKazantsev/Chatter/api/internal/adapters/handlers"
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	User *handlers.User
}

func NewController(clients clients.Clients) *Controller {
	return &Controller{
		User: handlers.NewUser(clients.UserClient),
	}
}

func InitRoutes(app *fiber.App, ctrl *Controller) {
	auth := app.Group("/user")
	auth.Post("/register", ctrl.User.Register)
	auth.Put("/login", ctrl.User.Login)
	auth.Put("/recovery", ctrl.User.PasswordRecovery)
	auth.Get("/email/verify", ctrl.User.VerifyCode)
	auth.Get("/email/send", ctrl.User.SendCode)
	auth.Get("/refresh", ctrl.User.UpdateTokens)

	fs := app.Group("/friends").Use(ctrl.User.ParseToken)
	fs.Put("/add", ctrl.User.SuggestFriendShip)
	fs.Delete("/refuse", ctrl.User.RefuseFriendShip)
}
