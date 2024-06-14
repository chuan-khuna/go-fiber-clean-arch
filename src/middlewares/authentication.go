package middlewares

import (
	"fiber-app/src/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// check if Authorization header is empty
	if authHeader == "" {

		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// check bearer token
	if !strings.HasPrefix(authHeader, "Bearer ") {

		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Invalid token format, must be Bearer"})
	}

	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	_, err := jwt.VerifyToken(tokenString)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Next()

}
