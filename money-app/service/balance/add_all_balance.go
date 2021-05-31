/*
残高に関わるアプリケーションロジック
*/
package balance_service

import (
	"github.com/jinzhu/gorm"
	"money-app/db"
	"money-app/repository"
	"sync"
)

type AddAllBalanceService struct {
	BalanceRepository           repository.BalanceRepository
	IdempotentCheckerRepository repository.IdempotentCheckerRepository
}

// デバッグ用に少なめに設定
const maxGettableCountAtOnce = 100

func (a AddAllBalanceService) Exec(amount uint, idempotentKey string) error {
	db := db.GetDB()
	return db.Transaction(func(tx *gorm.DB) error {
		idempotentChecker := a.IdempotentCheckerRepository.FindByIdempotentKey(idempotentKey)

		// 冪等keyが存在していたら後続の処理は実行しない
		if idempotentChecker.ID != 0 {
			return nil
		} else {
			err := a.IdempotentCheckerRepository.SaveTx(tx, idempotentKey)
			if err != nil {
				tx.Rollback()
			}
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
				err := a.BalanceRepository.UpdateBalancesTx(tx, ids, int(amount))
				if err != nil {
					tx.Rollback()
				}
			}()

			offset += maxGettableCountAtOnce
		}

		wg.Wait()
		return nil
	})
}
