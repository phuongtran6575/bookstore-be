package data

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migration
	db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Author{},
		&models.User{},
		&models.Address{},
		&models.Order{},
		&models.OrderItem{},
		&models.Admin{},
		&models.Role{},
		&models.Permission{},
		&models.AdminRole{},
		&models.RolePermission{},
	)

	DB = db
	return db
}
