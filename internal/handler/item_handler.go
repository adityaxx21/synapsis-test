package handler

import (
	"synapsis-backend-test/internal/service"
	"synapsis-backend-test/internal/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ItemAdminRoutes(api fiber.Router) {
	api.Get("/items/:id", DetailItem)
	api.Post("/items", CreateItem)
	api.Put("/items/:id", UpdateItem)
	api.Delete("/items/:id", DeleteItem)
}

func ItemAllRoutes(api fiber.Router)  {
	api.Get("/items", ListItem)
}

func CreateItem(c *fiber.Ctx) error {
	var item domain.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	if err := service.CreateItem(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := fiber.Map{
		"data":  item,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func UpdateItem(c *fiber.Ctx) error  {
	var req domain.Item
	var item *domain.Item

	id := c.Params("id")

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

    item, err := service.UpdateItem(id, &req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	response := fiber.Map{
		"data":  item,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item domain.Item

	if err := service.DeleteItem(id, &item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func DetailItem(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := service.DetailItem(id)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	response := fiber.Map{
		"data":  item,
	}


	return c.Status(fiber.StatusCreated).JSON(response)
}

func ListItem(c *fiber.Ctx) error {
	category := c.Query("category")
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}

	items, err := service.ListItem(category, page, limit)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	response := fiber.Map{
		"data":  items,
		"page":  page,
		"limit": limit,
	}

	return c.JSON(response)

}