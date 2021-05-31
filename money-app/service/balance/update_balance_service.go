package balance_service

import (
	"github.com/jinzhu/gorm"
	"money-app/db"
	"money-app/repository"
)

// Goにはクラスがないので構造体で定義
// type HogeがそのままHogeというクラスになる
type UpdateBalanceService struct {
	// 使用したい外部クラスを定義
	BalanceRepository           repository.BalanceRepository
	IdempotentCheckerRepository repository.IdempotentCheckerRepository
}

// UpdateBalanceServiceクラスに属するクラスメソッドとして定義
// 先頭の1文字をエイリアスとして定義するのが通常らしい
// (UpdateBalanceService => u)
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
