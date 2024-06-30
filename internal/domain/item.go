package domain

import (
	"time"

	"gorm.io/gorm"
	pq "github.com/lib/pq"
)

type Item struct {
	ID        		uint          	`gorm:"primaryKey";`
	Title	  		string      	`gorm:"not null"  validate:"required"`
	Description	  	string        	`gorm:"type:text;not null" validate:"required"` 
	Category	  	string        	`gorm:"type:text;not null" validate:"required"` 
	Price 			int				`gorm:"not null" validate:"required"`
	Size			pq.Float32Array `gorm:"type:float[]" validate:"required"`
	Weight			float32			`gorm:"not null" validate:"required"`
	Stock			int				`gorm:"not null" validate:"required,min=1,max=20"`
	Transactions 	[]Transaction `gorm:"many2many:transaction_items;"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
}

func FilterByCategory(category string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if category != "" {
			return db.Where("category = ?", category)
		}
		return db
	}
}