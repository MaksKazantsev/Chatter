package adapters

import (
	"github.com/MaksKazantsev/Chatter/api/internal/adapters/handlers"
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	User     *handlers.User
	Messages *handlers.Messages
	Files    *handlers.Files
}

func NewController(clients clients.Clients) *Controller {
	return &Controller{
		User:     handlers.NewUser(clients.UserClient),
		Messages: handlers.NewMessages(clients.MessagesClient),
		Files:    handlers.NewFiles(clients.FilesClient),
	}
}

func InitRoutes(app *fiber.App, ctrl *Controller) {
	app.Get("/chat/ws/join", websocket.New(ctrl.Messages.Join))
	app.Post("/files/upload", ctrl.Files.Upload)

	auth := app.Group("/auth")
	auth.Post("/register", ctrl.User.Register)
	auth.Put("/login", ctrl.User.Login)
	auth.Put("/recovery", ctrl.User.PasswordRecovery)
	auth.Get("/email/verify", ctrl.User.VerifyCode)
	auth.Get("/email/send", ctrl.User.SendCode)
	auth.Get("/refresh", ctrl.User.UpdateTokens)

	ch := app.Group("/chat")
	ch.Get("/message/:id", ctrl.Messages.DeleteMessage)
	ch.Get("/history/:receiverID", ctrl.Messages.GetHistory)
}
