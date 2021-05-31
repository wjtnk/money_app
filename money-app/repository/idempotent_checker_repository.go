package repository

import (
	"github.com/jinzhu/gorm"
	"money-app/db"
	"money-app/domain"
)

type IdempotentCheckerRepository struct {
	IdempotentChecker domain.IdempotentChecker
}

func (i *IdempotentCheckerRepository) FindByIdempotentKey(idempotentKey string) domain.IdempotentChecker {
	idempotentChecker := domain.IdempotentChecker{}
	db := db.GetDB()
	db.Where("idempotent_key = ?", idempotentKey).First(&idempotentChecker)
	return idempotentChecker
}

// SaveTx トランザクション使用時の保存
func (i *IdempotentCheckerRepository) SaveTx(tx *gorm.DB, idempotentKey string) error {
	idempotentChecker := domain.IdempotentChecker{IdempotentKey: idempotentKey}
	tx.NewRecord(idempotentChecker)
	return tx.Create(&idempotentChecker).Error
}
