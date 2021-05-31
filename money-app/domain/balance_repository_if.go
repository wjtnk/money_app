package domain

import (
	"github.com/jinzhu/gorm"
)

type BalanceRepositoryIf interface {
	FindByUserId(userId string) (Balance, error)
	Create(amount int, userId string)
	UpdateTx(tx *gorm.DB, id uint, amount int)
	UpdateBalances(Ids []uint, amount int)
	GetIds(limit uint, offset uint) []uint
}
