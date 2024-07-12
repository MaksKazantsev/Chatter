package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/clients"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type User struct {
	cl clients.UserClient
}

func NewUser(cl clients.UserClient) *User {
	return &User{cl: cl}
}

// Register godoc
// @Summary Register
// @Description Register new user
// @Tags Auth
// @Produce json
// @Param input body models.SignupReq true "user model"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/register [post]
func (u *User) Register(c *fiber.Ctx) error {
	var body models.SignupReq
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	aToken, rToken, err := u.cl.Register(c.Context(), body)
	if err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

// Login godoc
// @Summary Login
// @Tags Auth
// @Produce json
// @Param input body models.LoginReq true "user credentials"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/login [put]
func (u *User) Login(c *fiber.Ctx) error {
	var body models.LoginReq
	if err := c.BodyParser(&body); err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	aToken, rToken, err := u.cl.Login(c.Context(), body)
	if err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

// SendCode godoc
// @Summary SendCode
// @Description Send code to user verification/password recovery
// @Tags Auth
// @Produce json
// @Param email query string true "user email"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/email/send [get]
func (u *User) SendCode(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return utils.NewError("query parameter 'email' - is required", utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	if err := u.cl.SendCode(c.Context(), email); err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}
	c.Status(http.StatusOK)
	return nil
}

// VerifyCode godoc
// @Summary VerifyCode
// @Description User email code verification
// @Tags Auth
// @Produce json
// @Param type query string true "verification type"
// @Param code query string true "code"
// @Param email query string true "user email"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/email/verify [get]
func (u *User) VerifyCode(c *fiber.Ctx) error {
	var body models.VerifyCodeReq
	t, cd, email := c.Query("type"), c.Query("code"), c.Query("email")

	if email == "" || cd == "" || t == " " {
		return utils.NewError("query parameters 'email,code,type' - is required", utils.ERR_CLIENT_INVALID_ARGUMENT)
	}

	body.Code = cd
	body.Type = t
	body.Email = email

	aToken, rToken, err := u.cl.VerifyCode(c.Context(), body)
	if err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}

	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}

// PasswordRecovery godoc
// @Summary PasswordRecovery
// @Description Reset user password, requires verified email by VerifyCode
// @Tags Auth
// @Produce json
// @Param input body models.RecoveryReq true "recovery request"
//
//	@Success        201 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/recovery [put]
func (u *User) PasswordRecovery(c *fiber.Ctx) error {
	var body models.RecoveryReq
	if err := c.BodyParser(&body); err != nil {
		return utils.NewError(err.Error(), utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	if err := u.cl.PasswordRecovery(c.Context(), body); err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}
	c.Status(http.StatusOK)
	return nil
}

// UpdateTokens godoc
// @Summary UpdateTokens
// @Description Update user's token pair
// @Tags Auth
// @Produce json
// @Param Authorization header string true "token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/refresh [get]
func (u *User) UpdateTokens(c *fiber.Ctx) error {
	refresh := parseAuthHeader(c)
	if refresh == "" {
		return utils.NewError("refresh token not provided", utils.ERR_CLIENT_INVALID_ARGUMENT)
	}
	aToken, rToken, err := u.cl.UpdateTokens(c.Context(), refresh)
	if err != nil {
		if err != nil {
			st, msg := utils.HandleError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
	}
	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"token": aToken, "refreshToken": rToken})
	return nil
}
