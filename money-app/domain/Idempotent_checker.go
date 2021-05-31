package domain

import "github.com/jinzhu/gorm"

type IdempotentChecker struct {
	gorm.Model
	IdempotentKey string `gorm:"size:36;not null;unique"`
}
