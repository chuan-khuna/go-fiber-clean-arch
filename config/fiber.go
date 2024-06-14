package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: 2 * time.Second,
	}
}

func GetLoggerConfig() logger.Config {
	return logger.Config{
		Format:     "${time} | ${latency} | ${ip} | ${method} | ${path} | ${status} | ${error}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
	}
}
