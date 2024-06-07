package handlers

import "github.com/MaksKazantsev/SSO/api/internal/clients"

type Auth struct {
	cl clients.UserAuth
}

func NewAuth(cl clients.UserAuth) *Auth {
	return &Auth{cl: cl}
}

func (c *fiber.Ctx) Register() error {

}
