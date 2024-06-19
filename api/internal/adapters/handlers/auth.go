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
	aToken, rToken, err := a.cl.Register(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (a *Auth) Login(c *fiber.Ctx) error {
	var body models.LoginReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	aToken, rToken, err := a.cl.Login(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (a *Auth) SendCode(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		_ = c.Status(http.StatusBadGateway).SendString("query parameter 'email' - is required")
	}
	if err := a.cl.SendCode(c.Context(), email); err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	c.Status(http.StatusOK)
	return nil
}

func (a *Auth) VerifyCode(c *fiber.Ctx) error {
	var body models.VerifyCodeReq
	t, cd, email := c.Query("type"), c.Query("code"), c.Query("email")

	if email == "" || cd == "" || t == " " {
		_ = c.Status(http.StatusBadGateway).SendString("query parameters 'email,code,type' - is required")
	}

	body.Code = cd
	body.Type = t
	body.Email = email

	aToken, rToken, err := a.cl.VerifyCode(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (a *Auth) PasswordRecovery(c *fiber.Ctx) error {
	var body models.RecoveryReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	if err := a.cl.PasswordRecovery(c.Context(), body); err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	c.Status(http.StatusOK)
	return nil
}
