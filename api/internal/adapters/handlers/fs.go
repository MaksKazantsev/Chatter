package handlers

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (u *User) SuggestFriendShip(c *fiber.Ctx) error {
	token := parseAuthHeader(c)
	var req models.FriendShipReq
	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	req.Token = token

	if err := u.cl.SuggestFriendShip(c.Context(), req); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}

func (u *User) RefuseFriendShip(c *fiber.Ctx) error {

	token := parseAuthHeader(c)
	var req models.RefuseFriendShipReq
	if err := c.BodyParser(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	req.Token = token

	if err := u.cl.RefuseFriendShip(c.Context(), req); err != nil {
		st, msg := utils.HandleError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}
