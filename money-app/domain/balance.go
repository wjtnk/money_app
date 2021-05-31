package domain

import "github.com/jinzhu/gorm"

type Balance struct {
	gorm.Model
	Amount             int    `gorm:"not null"`
	UserId             string `gorm:"size:36;not null;unique"`
	IdempotentCheckers []IdempotentChecker
}

// UpdateBalance 残高を更新する。0以下にはならない
func (b *Balance) UpdateBalance(amount int) {
	b.Amount = b.Amount + amount
	if b.Amount < 0 {
		b.Amount = 0
	}
}
