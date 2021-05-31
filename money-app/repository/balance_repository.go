/*
レポジトリ
*/
package repository

import (
	"github.com/jinzhu/gorm"
	"money-app/db"
	"money-app/domain"
)

type BalanceRepository struct {
}

func (b BalanceRepository) FindByUserId(userId string) (domain.Balance, error) {
	balance := domain.Balance{}
	db := db.GetDB()
	err := db.Where("user_id = ?", userId).First(&balance).Error

	if err != nil {
		return balance, err
	}

	return balance, nil
}

func (b BalanceRepository) Create(amount int, userId string) {
	balance := domain.Balance{Amount: amount, UserId: userId}
	db := db.GetDB()
	db.Create(&balance)
}

// UpdateTx トランザクション使用時の更新
func (b BalanceRepository) UpdateTx(tx *gorm.DB, id uint, amount int) {
	balance := domain.Balance{}
	tx.Where("id = ?", id).First(&balance).UpdateColumn("amount", amount)
}

func (b BalanceRepository) UpdateBalances(Ids []uint, amount int) {
	balance := domain.Balance{}
	db := db.GetDB()
	db.Model(&balance).Where("id IN (?)", Ids).Updates(
		map[string]interface{}{
			"amount": gorm.Expr("amount + ?", amount),
		})
}

func (b BalanceRepository) GetIds(limit uint, offset uint) []uint {
	var ids []uint
	var balances []domain.Balance

	db := db.GetDB()
	db.Limit(limit).Offset(offset).Find(&balances).Pluck("id", &ids)
	return ids
}
