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

func (b BalanceRepository) GetIds(limit uint, offset uint) []uint {
	var ids []uint
	var balances []domain.Balance

	db := db.GetDB()
	db.Limit(limit).Offset(offset).Find(&balances).Pluck("id", &ids)
	return ids
}

func (b BalanceRepository) Create(amount int, userId string) {
	balance := domain.Balance{Amount: amount, UserId: userId}
	db := db.GetDB()
	db.Create(&balance)
}

// UpdateTx トランザクション使用時の更新
func (b BalanceRepository) UpdateTx(tx *gorm.DB, id uint, amount int) error {
	balance := domain.Balance{}
	err := tx.Where("id = ?", id).First(&balance).UpdateColumn("amount", amount).Error

	if err != nil {
		return err
	}
	return nil
}

func (b BalanceRepository) UpdateBalancesTx(tx *gorm.DB, Ids []uint, amount int) error {
	balance := domain.Balance{}

	return tx.Model(&balance).Where("id IN (?)", Ids).Updates(
		map[string]interface{}{
			"amount": gorm.Expr("amount + ?", amount),
		}).Error
}