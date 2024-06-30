package handler

import (
	"strconv"

	"synapsis-backend-test/internal/service"
	"synapsis-backend-test/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CartRoutes(api fiber.Router) {
	api.Get("/carts", ListCart)
	api.Post("/carts", StoreCart)
	api.Put("/carts/:itemId", UpdateCart)
	api.Delete("/carts/:itemId", DeleteCart)
}

func StoreCart(c *fiber.Ctx) error {
	var cart domain.Cart

	jwt_user := c.Locals("user").(*jwt.Token)
	claims := jwt_user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

    user, err := service.DetailUser(username)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, errs := service.StoreCart(user, &cart)

    if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }

	response := fiber.Map{
		"data":  result,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}



func ListCart(c *fiber.Ctx) error {
	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

    user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}

	carts, err := service.ListCart(user, page, limit)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
	
	response := fiber.Map{
		"data":  carts,
		"page":  page,
		"limit": limit,
	}

	return c.JSON(response)
}


func UpdateCart(c *fiber.Ctx) error  {
	var req domain.Cart
	var cart *domain.CartItem

	id := c.Params("itemId")

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

    user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

    cart, err := service.UpdateCart(id, user.ID, &req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	response := fiber.Map{
		"data":  cart,
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}


func DeleteCart(c *fiber.Ctx) error {
	id := c.Params("itemId")
	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

    user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }

	if err := service.DeleteCart(user.ID, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
