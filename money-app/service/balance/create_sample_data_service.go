package balance_service

import (
	"github.com/google/uuid"
	"money-app/repository"
	"sync"
)

type CreateSampleDataService struct {
	BalanceRepository repository.BalanceRepository
}

func (c CreateSampleDataService) Exec() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go c.createBalanceRecords(&wg)
	wg.Wait()
}

func (c CreateSampleDataService) createBalanceRecords(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		uuidV4 := uuid.New()
		c.BalanceRepository.Create(0, uuidV4.String())
	}
}
