package domain

import (
	"time"
	pq "github.com/lib/pq"

)

type (
	Cart struct {
		UserID  uint `gorm:"primary_key"`
		ItemID  uint `gorm:"primary_key"`
		Total   int  `gorm:"not null;int;" validate:"required"`  
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CartItem struct {
		ID        		uint          	`gorm:"primaryKey";`
		Title	  		string      	`gorm:"not null"  validate:"required"`
		Description	  	string        	`gorm:"type:text;not null" validate:"required"` 
		Category	  	string        	`gorm:"type:text;not null" validate:"required"` 
		Price 			int				`gorm:"not null" validate:"required"`
		Size			pq.Float32Array `gorm:"type:float[]" validate:"required"`
		Weight			float32			`gorm:"not null" validate:"required"`
		Stock			int				`gorm:"not null" validate:"required"`
		UserID  		uint 
		ItemID 			uint 
		Total			int				`gorm:"not null" validate:"required"`
	}
	
)