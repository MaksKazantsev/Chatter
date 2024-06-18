package handlers

import (
	"github.com/MaksKazantsev/SSO/api/internal/clients"
	"github.com/MaksKazantsev/SSO/api/internal/models"
	"github.com/MaksKazantsev/SSO/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Auth struct {
	cl clients.UserAuth
}

func NewAuth(cl clients.UserAuth) *Auth {
	return &Auth{cl: cl}
}

func (a *Auth) Register(c *fiber.Ctx) error {
	var body models.SignupReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	token, err := a.cl.Register(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": token})
	return nil
}

func (a *Auth) Login(c *fiber.Ctx) error {
	var body models.LoginReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	token, err := a.cl.Login(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": token})
	return nil
}

func (a *Auth) Reset(c *fiber.Ctx) error {
	var body models.ResetReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	if err := a.cl.Reset(c.Context(), body); err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	return nil
}
