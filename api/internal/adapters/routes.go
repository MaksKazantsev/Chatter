package adapters

import (
	"github.com/MaksKazantsev/Chatter/api/internal/adapters/handlers"
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	auth := app.Group("/auth")
	auth.Post("/register", ctrl.User.Register)
	auth.Put("/login", ctrl.User.Login)
	auth.Put("/recovery", ctrl.User.PasswordRecovery)
	auth.Get("/email/verify", ctrl.User.VerifyCode)
	auth.Get("/email/send", ctrl.User.SendCode)
	auth.Get("/refresh", ctrl.User.UpdateTokens)

	ch := app.Group("/chat")
	ch.Delete("/message/:id", ctrl.Messages.DeleteMessage)
	ch.Get("/history/:targetID", ctrl.Messages.GetHistory)

	user := app.Group("/user")
	user.Get("/friends/get", ctrl.User.GetFriendsSection)
	user.Delete("/friends/delete/:targetID", ctrl.User.DeleteFriend)
	user.Post("/friends/suggest/:targetID", ctrl.User.SuggestFs)
	user.Get("/friends/refuse/:targetID", ctrl.User.RefuseFs)
	user.Get("/friends/accept/:targetID", ctrl.User.AcceptFs)

	user.Put("/profile/edit", ctrl.User.EditProfile)
	user.Put("/profile/avatar/edit", ctrl.User.EditAvatar)
	user.Delete("/profile/avatar/delete", ctrl.User.DeleteAvatar)
	user.Get("/profile/:targetID", ctrl.User.GetProfile)
}
