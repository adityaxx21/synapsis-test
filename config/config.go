package config

import (
	"synapsis-backend-test/internal/domain"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var DB *gorm.DB
var Mid snap.Client



func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=synapsisDb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	fmt.Println("Database connection successful")
}


func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{}, 
		&domain.Item{}, 
		&domain.Cart{}, 
		&domain.Transaction{},
		&domain.TransactionItem{},
	)
}

func InitializeSnapClient() {
	Mid.New("SB-Mid-server-3CD8qhM1_29NjzaDByb_6FTs", midtrans.Sandbox)
}
