package models

import (
	"time"
)

type Review struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	ProductID   uint          `gorm:"not null" json:"product_id"`
	UserID      uint          `gorm:"not null" json:"user_id"`
	OrderItemID uint          `gorm:"not null" json:"order_item_id"`
	Rating      int8          `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Title       string        `gorm:"size:300" json:"title"`
	Body        string        `gorm:"type:text" json:"body"`
	IsApproved  bool          `gorm:"not null;default:false" json:"is_approved"`
	ApprovedAt  *time.Time     `json:"approved_at"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	Images      []ReviewImage `gorm:"foreignKey:ReviewID" json:"images"`
}

type ReviewImage struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ReviewID  uint   `gorm:"not null" json:"review_id"`
	URL       string `gorm:"size:500;not null" json:"url"`
	SortOrder int8   `json:"sort_order"`
}

type ReviewHelpful struct {
	ReviewID  uint      `gorm:"primaryKey" json:"review_id"`
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	IsHelpful bool      `gorm:"not null" json:"is_helpful"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
