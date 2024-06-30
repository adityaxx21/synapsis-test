package main

import (
	"synapsis-backend-test/config"
	"synapsis-backend-test/internal/handler"
	"synapsis-backend-test/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()
	config.ConnectDB()
	config.Migrate(config.DB)
	config.InitializeSnapClient()
	api := app.Group("/api")

	// handle unautorize route
	handler.UserRoutes(api)

	// handle authorize route 
	api.Use(middleware.JWTMiddleware)
	handler.SimpleRoutes(api)
	handler.CartRoutes(api)
	handler.TransactionRoutes(api)
	handler.ItemAllRoutes(api)

	// handle admin authorize route
	api.Use(middleware.JWTMiddleware, middleware.AdminMiddleware)
	handler.ItemAdminRoutes(api)

	app.Listen(":3000")
}