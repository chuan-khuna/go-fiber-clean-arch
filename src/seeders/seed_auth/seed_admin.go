package seed_auth

import (
	auth_entities "fiber-app/src/entities/auth"
	"fiber-app/src/orm"
	"fiber-app/src/utils/password"
)

func SeedAdmins() {
	db := orm.InitDB()

	hashedPassword, _ := password.HashPassword("1234")

	admin := auth_entities.AdminAccount{
		Email:     "admin@admin.com",
		Password:  hashedPassword,
		IsRevoked: false,
	}

	db.FirstOrCreate(&admin)
}
