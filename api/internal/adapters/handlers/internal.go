package handlers

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func parseAuthHeader(c *fiber.Ctx) string {
	auth := c.Get("Authorization")

	vals := strings.Split(auth, " ")
	if len(vals) != 2 {
		return ""
	}
	return vals[1]
}

func parseWSAuthHeader(c *websocket.Conn) string {
	auth := c.Headers("Authorization", "")
	if auth == "" {
		return auth
	}
	vals := strings.Split(auth, " ")
	if len(vals) != 2 {
		return ""
	}
	return vals[1]
}
