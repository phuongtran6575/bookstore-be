package models

import (
	"time"
)

type CouponType string

const (
	CouponTypePercent CouponType = "percent"
	CouponTypeFixed   CouponType = "fixed"
)

type Coupon struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	Code              string     `gorm:"size:50;uniqueIndex;not null" json:"code"`
	Description       string     `gorm:"type:text" json:"description"`
	Type              CouponType `gorm:"type:varchar(20);not null" json:"type"` // percent, fixed
	Value             float64    `gorm:"type:decimal(10,2);not null" json:"value"`
	MinOrderValue     *float64  `gorm:"type:decimal(12,2)" json:"min_order_value"`
	MaxDiscountAmount *float64  `gorm:"type:decimal(12,2)" json:"max_discount_amount"`
	MaxUses           *int      `json:"max_uses"`
	MaxUsesPerUser    *int      `json:"max_uses_per_user"`
	UsedCount         int       `gorm:"not null;default:0" json:"used_count"`
	IsActive          bool      `gorm:"not null;default:true" json:"is_active"`
	StartsAt          *time.Time `json:"starts_at"`
	ExpiresAt         *time.Time `json:"expires_at"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type CouponUsage struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CouponID        uint      `gorm:"not null" json:"coupon_id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	OrderID         uint      `gorm:"not null" json:"order_id"`
	DiscountApplied float64   `gorm:"type:decimal(12,2);not null" json:"discount_applied"`
	UsedAt          time.Time `gorm:"not null;autoCreateTime" json:"used_at"`
}
