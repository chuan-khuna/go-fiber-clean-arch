package usecases

import "fiber-app/src/entities"

type PingRepository interface {
	// use case of ping repository

	Log(p entities.Ping) error
}
