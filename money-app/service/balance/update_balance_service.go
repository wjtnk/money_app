package balance_service

import (
	"github.com/jinzhu/gorm"
	"money-app/db"
	"money-app/repository"
)

type UpdateBalanceService struct {
	BalanceRepository           repository.BalanceRepository
	IdempotentCheckerRepository repository.IdempotentCheckerRepository
}

func (u UpdateBalanceService) Exec(userId string, amount int, idempotentKey string) error {
	db := db.GetDB()
	return db.Transaction(func(tx *gorm.DB) error {
		idempotentChecker := u.IdempotentCheckerRepository.FindByIdempotentKey(idempotentKey)

		// 冪等keyが存在していたら後続の処理は実行しない
		if idempotentChecker.ID != 0 {
			return nil
		} else {
			u.IdempotentCheckerRepository.SaveTx(tx, idempotentKey)
		}

		balance, err := u.BalanceRepository.FindByUserId(userId)

		if err != nil {
			return err
		}

		balance.UpdateBalance(amount)

		u.BalanceRepository.UpdateTx(tx, balance.ID, balance.Amount)
		return nil
	})
}
