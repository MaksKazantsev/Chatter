package middleware

import (
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func WithErrorWrapper(handler fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := handler(ctx)

		if err != nil {
			st, msg := utils.HandleError(err)
			_ = ctx.JSON(fiber.Map{"status": st, "error": msg})
		}
		return nil
	}
}
