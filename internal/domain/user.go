package domain

import (
	"time"

	"gorm.io/gorm"
)

type (
	 User struct {
		ID        uint           `gorm:"primaryKey"`
		Username 	string 		 `validate:"required,min=3,max=32"`
		Password 	string 		 `validate:"required,min=6"`
		Role     	string 		 `validate:"required,oneof=admin user"`
		FName	 	string		 `validate:"required"`
		LName 		string		 `validate:"required"`
		Phone		string		 `validate:"required"`
		Address	 	string		 `validate:"required"`
		City	 	string		 `validate:"required"`
		Postcode 	string		 `validate:"required"`
		CreatedAt 	time.Time
		UpdatedAt 	time.Time
		DeletedAt 	gorm.DeletedAt `gorm:"index"`
	}
	
	
	UserJwt struct {
		username	string	`validate:"required"`
		role		string	`validate:"required"`
		exp			int		`validate:"required"`
	}
)
