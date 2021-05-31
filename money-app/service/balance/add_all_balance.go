/*
残高に関わるアプリケーションロジック
*/
package balance_service

import (
	"money-app/repository"
	"sync"
)

type AddAllBalanceService struct {
	BalanceRepository           repository.BalanceRepository
	IdempotentCheckerRepository repository.IdempotentCheckerRepository
}

// デバッグ用に少なめに設定
const maxGettableCountAtOnce = 100

func (a AddAllBalanceService) Exec(amount uint, idempotentKey string) {
	idempotentChecker := a.IdempotentCheckerRepository.FindByIdempotentKey(idempotentKey)

	// 冪等keyが存在していたら後続の処理は実行しない
	if idempotentChecker.ID != 0 {
		return
	} else {
		a.IdempotentCheckerRepository.Save(idempotentKey)
	}

	var wg sync.WaitGroup
	offset := 0

	for {
		ids := a.BalanceRepository.GetIds(maxGettableCountAtOnce, uint(offset))
		if len(ids) == 0 {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			a.BalanceRepository.UpdateBalances(ids, int(amount))
		}()

		offset += maxGettableCountAtOnce
	}

	wg.Wait()
}
