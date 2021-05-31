package domain

import (
	"github.com/jinzhu/gorm"
)

type IdempotentCheckerRepositoryIf interface {
	FindByIdempotentKey(idempotentKey string) IdempotentChecker
	Save(idempotentKey string)
	SaveTx(tx *gorm.DB, idempotentKey string)
}
