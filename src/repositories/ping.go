package repositories

import (
	"fiber-app/src/entities"
	"fiber-app/src/usecases"
	"fmt"

	"gorm.io/gorm"
)

type GormPingRepository struct {
	db *gorm.DB
}

func NewGormPingRepository(db *gorm.DB) usecases.PingRepository {
	return &GormPingRepository{db: db}
}

func (r *GormPingRepository) Log(p entities.Ping) error {
	fmt.Println("Gorm Repo: got ping message: ", p.Message)
	// save to db, but for now just print it
	return nil
}
