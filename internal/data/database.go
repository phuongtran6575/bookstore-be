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
	//sslmode := os.Getenv("DB_SSLMODE")
	//timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migration
	db.AutoMigrate(
		// Catalog
		&models.ProductType{},
		&models.Product{},
		&models.ProductImage{},
		&models.Category{},
		&models.ProductCategory{},
		&models.Author{},
		&models.ProductAuthor{},
		&models.AttributeDefinition{},
		&models.AttributeSelectOption{},
		&models.ProductAttributeValue{},
		&models.VariantOptionType{},
		&models.VariantOptionValue{},
		&models.ProductVariant{},
		&models.VariantOptionAssignment{},
		&models.Bundle{},
		&models.BundleItem{},

		// Users
		&models.User{},
		&models.Address{},
		&models.UserSession{},
		&models.Cart{},
		&models.CartItem{},
		&models.Wishlist{},

		// Orders
		&models.Order{},
		&models.OrderItem{},
		&models.OrderStatusLog{},

		// Payments
		&models.PaymentMethod{},
		&models.Payment{},
		&models.Refund{},
		&models.StoreWallet{},
		&models.WalletTransaction{},

		// Shipping
		&models.Carrier{},
		&models.ShippingMethod{},
		&models.ShippingZone{},
		&models.ShippingRate{},
		&models.Shipment{},
		&models.ShipmentTrackingEvent{},

		// Promotions
		&models.Coupon{},
		&models.CouponUsage{},

		// Reviews
		&models.Review{},
		&models.ReviewImage{},
		&models.ReviewHelpful{},

		// Inventory & Loyalty
		&models.LoyaltyAccount{},
		&models.LoyaltyTransaction{},
		&models.StoreLocation{},
		&models.LocationInventory{},
		&models.InventoryLog{},

		// RBAC
		&models.Admin{},
		&models.Role{},
		&models.Permission{},
		&models.AdminRole{},
		&models.RolePermission{},

		// System
		&models.AuditLog{},
		&models.Notification{},
		&models.Setting{},
	)

	DB = db
	return db
}
