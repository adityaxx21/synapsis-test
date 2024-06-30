package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	jwtware "github.com/gofiber/jwt/v3"
)

var secretKey = []byte("secret")

func JWTMiddleware(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			c.Locals("username", claims["username"].(string))
			c.Locals("role", claims["role"].(string))
			return c.Next()
		},
	})(c)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "missing or malformed JWT") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing or malformed JWT"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
	}
	return nil
}

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}
	return c.Next()
}