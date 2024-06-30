package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type User struct {
	cl clients.User
}

func NewUser(cl clients.User) *User {
	return &User{cl: cl}
}

func (u *User) Register(c *fiber.Ctx) error {
	var body models.SignupReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	aToken, rToken, err := u.cl.Register(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (u *User) Login(c *fiber.Ctx) error {
	var body models.LoginReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	aToken, rToken, err := u.cl.Login(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (u *User) SendCode(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		_ = c.Status(http.StatusBadGateway).SendString("query parameter 'email' - is required")
	}
	if err := u.cl.SendCode(c.Context(), email); err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	c.Status(http.StatusOK)
	return nil
}

func (u *User) VerifyCode(c *fiber.Ctx) error {
	var body models.VerifyCodeReq
	t, cd, email := c.Query("type"), c.Query("code"), c.Query("email")

	if email == "" || cd == "" || t == " " {
		_ = c.Status(http.StatusBadGateway).SendString("query parameters 'email,code,type' - is required")
	}

	body.Code = cd
	body.Type = t
	body.Email = email

	aToken, rToken, err := u.cl.VerifyCode(c.Context(), body)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (u *User) PasswordRecovery(c *fiber.Ctx) error {
	var body models.RecoveryReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	if err := u.cl.PasswordRecovery(c.Context(), body); err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	c.Status(http.StatusOK)
	return nil
}

func (u *User) UpdateTokens(c *fiber.Ctx) error {
	refresh := parseAuthHeader(c)
	if refresh == "" {
		_ = c.Status(http.StatusBadRequest).SendString("no refresh token found")
		return nil
	}
	aToken, rToken, err := u.cl.UpdateTokens(c.Context(), refresh)
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

func (u *User) ParseToken(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	vals := strings.Split(auth, ",")
	if len(vals) != 2 {
		c.Status(http.StatusMethodNotAllowed).SendString("token is not provided")
		return nil
	}
	_, err := u.cl.ParseToken(c.Context(), vals[1])
	if err != nil {
		code, msg := utils.HandleError(err)
		_ = c.Status(code).SendString(msg)
		return nil
	}
	return nil
}
