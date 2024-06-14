package routers

import (
	auth_controllers "fiber-app/src/controllers/auth"
	"fiber-app/src/orm"

	// auth_entities "fiber-app/src/entities/auth"
	auth_repositories "fiber-app/src/repositories/auth"
	auth_usecases "fiber-app/src/usecases/auth"

	"github.com/gofiber/fiber/v2"
)

func AdminAuthRouter(a *fiber.App) {
	db := orm.InitDB()

	adminAuthRepo := auth_repositories.NewGormAdminAuthRepository(db)
	adminAuthUseCase := auth_usecases.NewAdminAuthService(adminAuthRepo)
	adminAuthController := auth_controllers.NewAuthController(adminAuthUseCase)

	// routerGroup := a.Group("/admin")
	// routerGroup.Post("/login", adminAuthController.Login)

	a.Post("/login", adminAuthController.Login)
	a.Post("/refresh", adminAuthController.RefreshAccessToken)
	a.Post("/logout", adminAuthController.Logout)
	a.Post("/verify", adminAuthController.VerifyToken)
}
