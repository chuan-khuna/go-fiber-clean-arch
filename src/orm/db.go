package orm

import (
	"fiber-app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dbHost := config.DBHost
	dbPort := config.DBPort
	dbUser := config.DBUser
	dbPassword := config.DBPassword
	dbName := config.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("🔥 failed to connect database")
	}

	return db
}

func RunAutoMigrate(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		fmt.Println("Auto migrating model", model)
		db.AutoMigrate(model)
	}
}
