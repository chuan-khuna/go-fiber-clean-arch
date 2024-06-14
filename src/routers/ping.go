package routers

import (
	"fiber-app/src/controllers"
	"fiber-app/src/orm"
	"fiber-app/src/repositories"
	"fiber-app/src/usecases"

	"github.com/gofiber/fiber/v2"
)

func PingRouter(a *fiber.App) {
	a.Get("/ping", controllers.Ping)

	db := orm.InitDB()
	pingRepo := repositories.NewGormPingRepository(db)
	PingService := usecases.NewPingService(pingRepo)

	// ping controller called at the endpoint
	pingController := controllers.NewPingController(PingService)
	a.Post("/ping", pingController.Log)
}
