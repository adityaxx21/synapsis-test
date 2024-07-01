package config

import (
	"synapsis-backend-test/internal/domain"
	"github.com/midtrans/midtrans-go/example"
	"os"
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
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=0.0.0.0 sslmode=disable user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", user, password, dbname, port)
	fmt.Println(dsn)
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
	Mid.New(example.SandboxServerKey1, midtrans.Sandbox)
}
