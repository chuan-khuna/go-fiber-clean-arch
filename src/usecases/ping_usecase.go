package usecases

import (
	"fiber-app/src/entities"
)

type PingUseCase interface {
	// business logic of ping
	// what ping can do?

	// log: ping can be saved to db
	// but for mocking purpose, it just return the message
	Log(p entities.Ping) error
}

type PingService struct {
	repo PingRepository
}

func NewPingService(repo PingRepository) PingUseCase {
	// ping service is a use case
	// so it should implement methods in PingUseCase
	return &PingService{repo: repo}
}

func (s *PingService) Log(p entities.Ping) error {
	// fmt.Println("PingService p.Message: ", p.Message)
	err := s.repo.Log(p)
	if err != nil {
		return err
	}
	return nil
}
