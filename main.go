package main

import (
	"fiber-app/config"
	auth_entities "fiber-app/src/entities/auth"
	"fiber-app/src/orm"
	"fiber-app/src/routers"
	auth_routers "fiber-app/src/routers/auth"
	"fiber-app/src/seeders"
	seed_auth "fiber-app/src/seeders/seed_auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	fiberConfig := config.GetFiberConfig()

	app := fiber.New(fiberConfig)

	app.Use(logger.New(
		config.GetLoggerConfig(),
	))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	models := []interface{}{
		&auth_entities.AdminAccount{},
		&auth_entities.AdminRefreshToken{},
	}

	db := orm.InitDB()
	orm.RunAutoMigrate(db, models)

	seeders.RunSeed(
		seed_auth.SeedAdmins,
	)

	// register routers
	routers.PingRouter(app)
	auth_routers.AdminAuthRouter(app)

	PORT := config.Port
	app.Listen(":" + PORT)
}
