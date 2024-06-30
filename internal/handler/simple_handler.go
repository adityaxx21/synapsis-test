package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SimpleRoutes(api fiber.Router) {
	api.Get("/check", Check)
}

func Check(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Status":"authorized"})
}