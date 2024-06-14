package auth_controllers

import (
	auth_entities "fiber-app/src/entities/auth"
	auth_usecases "fiber-app/src/usecases/auth"

	"github.com/gofiber/fiber/v2"
)

type AdminAuthController struct {
	authUseCase auth_usecases.AdminAuthUseCase
}

func NewAuthController(authUseCase auth_usecases.AdminAuthUseCase) *AdminAuthController {
	return &AdminAuthController{authUseCase: authUseCase}
}

func (h *AdminAuthController) Login(c *fiber.Ctx) error {
	payload := new(auth_entities.LoginPayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, tokenPair, err := h.authUseCase.Login(*payload)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": tokenPair,
	})

}

func (h *AdminAuthController) RefreshAccessToken(c *fiber.Ctx) error {
	// get new access token-refresh token pair

	refreshTokenString := c.FormValue("refreshToken")

	tokenPair, err := h.authUseCase.RefreshAccessToken(refreshTokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": tokenPair,
	})
}

func (h *AdminAuthController) Logout(c *fiber.Ctx) error {
	tokenPair := auth_entities.TokenPair{
		AccessToken:  c.FormValue("accessToken"),
		RefreshToken: c.FormValue("refreshToken"),
	}

	err := h.authUseCase.Logout(tokenPair)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *AdminAuthController) VerifyToken(c *fiber.Ctx) error {
	tokenString := c.FormValue("token")

	err := h.authUseCase.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Token is valid"})
}
