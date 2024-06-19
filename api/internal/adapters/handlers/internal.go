package handlers

import (
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
