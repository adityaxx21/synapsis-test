package handler

import (
	"synapsis-backend-test/internal/service"
	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/internal/validation"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Post("/login", Login)
	api.Post("/register", Register)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := service.Login(data["username"], data["password"])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func Register(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validation.ValidateStruct(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.Register(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
