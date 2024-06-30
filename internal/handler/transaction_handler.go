package handler

import (
	"fmt"
	"strconv"
	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/internal/service"
	"synapsis-backend-test/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"


)

func TransactionRoutes(api fiber.Router) {
	api.Get("/orders/:id", DetailOrder)
	api.Get("/orders", ListOrder)
	api.Post("/orders", CreateOrder)
	api.Post("/order-carts", CreateOrderCart)
	api.Post("/midtrans", CreateTransaction)
	// api.Put("/items/:id", UpdateItem)
	// api.Delete("/items/:id", DeleteItem)
}

func CreateOrder(c *fiber.Ctx) error  {
	var transaction domain.TransactionRequest
	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }
	
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.CreateOrder(user, &transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)

}	

func CreateOrderCart(c *fiber.Ctx) error {
	var transaction domain.TransactionCartRequest
	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }
	
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.CreateOrderCart(user, &transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}


func DetailOrder(c *fiber.Ctx) error {
	var transaction *domain.Transaction

	id := c.Params("id")

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	user, errs := service.DetailUser(username)
	if errs != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errs.Error()})
    }

	transaction,  err := service.DetailOrder(user, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := fiber.Map{
		"data":  transaction,
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func ListOrder(c *fiber.Ctx) error  {
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

	orders, err := service.ListOrder(user, page, limit)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
	
	response := fiber.Map{
		"data":  orders,
		"page":  page,
		"limit": limit,
	}

	return c.JSON(response)
}


func CreateTransaction(c *fiber.Ctx) error {
	// Optional : here is how if you want to set append payment notification for this request
	config.Mid.Options.SetPaymentAppendNotification("http://locahost:3000/")
	// Send request to Midtrans Snap API

	resp, err := config.Mid.CreateTransaction(service.GenerateSnapReqq())
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)

	return c.SendStatus(fiber.StatusCreated)
}



