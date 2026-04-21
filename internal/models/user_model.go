package models

import (
	"time"
)

type User struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Username        string         `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email           string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash    string         `gorm:"size:255;not null" json:"-"`
	FullName        string         `gorm:"size:300;not null" json:"full_name"`
	Phone           string         `gorm:"size:20" json:"phone"`
	Avatar          string         `gorm:"size:500" json:"avatar"`
	IsActive        bool           `gorm:"not null;default:true" json:"is_active"`
	EmailVerified   bool           `gorm:"not null;default:false" json:"email_verified"`
	EmailVerifiedAt *time.Time      `json:"email_verified_at"`
	LastLoginAt     *time.Time      `json:"last_login_at"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Addresses       []Address      `gorm:"foreignKey:UserID" json:"addresses"`
	Wishlist        []Product      `gorm:"many2many:wishlists;" json:"wishlist"`
}

type Address struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	RecipientName string    `gorm:"size:300;not null" json:"recipient_name"`
	Phone         string    `gorm:"size:20;not null" json:"phone"`
	Address       string    `gorm:"size:500;not null" json:"address"`
	Ward          string    `gorm:"size:200" json:"ward"`
	District      string    `gorm:"size:200;not null" json:"district"`
	City          string    `gorm:"size:200;not null" json:"city"`
	Country       string    `gorm:"size:100;not null;default:VN" json:"country"`
	ZipCode       string    `gorm:"size:20" json:"zip_code"`
	IsDefault     bool      `gorm:"not null;default:false" json:"is_default"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserSession struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	TokenHash  string    `gorm:"size:255;uniqueIndex;not null" json:"-"`
	DeviceInfo string    `gorm:"size:500" json:"device_info"`
	IPAddress  string    `gorm:"size:45" json:"ip_address"`
	ExpiresAt  time.Time `gorm:"not null" json:"expires_at"`
	RevokedAt  *time.Time `json:"revoked_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Cart struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	UserID       *uint      `json:"user_id"`
	SessionToken string     `gorm:"size:255;uniqueIndex" json:"session_token"`
	ExpiresAt    *time.Time  `json:"expires_at"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Items        []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `gorm:"not null;uniqueIndex:idx_cart_product_variant" json:"cart_id"`
	ProductID uint    `gorm:"not null;uniqueIndex:idx_cart_product_variant" json:"product_id"`
	VariantID *uint   `gorm:"uniqueIndex:idx_cart_product_variant" json:"variant_id"`
	Quantity  int     `gorm:"not null;default:1" json:"quantity"`
}

type Wishlist struct {
	UserID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
}
