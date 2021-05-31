package db

import (
	"money-app/domain"
)

func Migrate() {
	db.AutoMigrate(&domain.Balance{})
	db.AutoMigrate(&domain.IdempotentChecker{})
}

func DropAllTable() {
	db.DropTable(&domain.Balance{})
	db.DropTable(&domain.IdempotentChecker{})
}
