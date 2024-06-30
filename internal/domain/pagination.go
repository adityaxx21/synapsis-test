package domain

import (
	"gorm.io/gorm"
)

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
