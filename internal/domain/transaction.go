package domain

import (
	"time"

	"gorm.io/gorm"
	// pq "github.com/lib/pq"
)

type (
	Transaction struct {
		ID                uint           `gorm:"primaryKey"`
		UserID            uint           `gorm:"not null"`
		GrossAmount       int       	 `gorm:"not null"`
		Status            string         `gorm:"type:varchar(100);not null;default:paid"`
		TransactionDate   time.Time      `gorm:"not null"`
		Items             []Item         `gorm:"many2many:transaction_items;"`
		OrderType		  string 		 `gorm:"not null;default:bank_transfer"`
		MidtransOrderID   string         `gorm:"type:varchar(100)"`
		MidtransStatus    string         `gorm:"type:varchar(100)"`
		MidtransInvoice	  string         `gorm:"type:text"`
		CreatedAt         time.Time
		UpdatedAt         time.Time
		DeletedAt         gorm.DeletedAt `gorm:"index"`
	}

	TransactionRequest struct {
		GrossAmount       int       	 `validate:"required"`
		Status            string         `validate:"required"`
		ItemID            int	         `validate:"required"`
		Total			  int			 `validate:"required"`
		OrderType		  string 		 `validate:"required"`
		MidtransOrderID   string         `gorm:"type:varchar(100)"`
		MidtransStatus    string         `gorm:"type:varchar(100)"`
		MidtransInvoice	  string         `gorm:"type:text"`
	}

	TransactionCartRequest struct {
		GrossAmount       int       	 	`validate:"required"`
		Status            string         	`validate:"required"`
		Items             []TransactionItem	`validate:"required"`
		OrderType		  string 		 	`validate:"required"`
		MidtransOrderID   string         	`gorm:"type:varchar(100)"`
		MidtransStatus    string         	`gorm:"type:varchar(100)"`
		MidtransInvoice	  string         	`gorm:"type:text"`
	}

	TransactionItem struct {
		TransactionID int `gorm:"primaryKey"`
		ItemID        int `gorm:"primaryKey"`
		Total		  int `gorm:"not null"`
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
)