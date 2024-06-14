package controllers

import (
	"fiber-app/src/entities"
	"fiber-app/src/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong", "timestamp": time.Now().Format("2006-01-02 15:04:05.000")})
}

type PingController struct {
	pingUseCase usecases.PingUseCase
}

func NewPingController(pingUseCase usecases.PingUseCase) *PingController {
	return &PingController{pingUseCase: pingUseCase}
}

func (h *PingController) Log(c *fiber.Ctx) error {
	// h stands for handler, ie controller
	ping := new(entities.Ping)

	if err := c.BodyParser(ping); err != nil {
		ping.Message = "no message"
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// call use case
	if err := h.pingUseCase.Log(*ping); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message_received": ping.Message, "timestamp": time.Now().Format("2006-01-02 15:04:05.000")})
}
