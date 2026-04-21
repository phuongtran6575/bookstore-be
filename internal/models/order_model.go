package models

import (
	"time"

	"gorm.io/datatypes"
)

type Order struct {
	ID             uint             `gorm:"primaryKey" json:"id"`
	UserID         uint             `gorm:"not null" json:"user_id"`
	AddressID      uint             `gorm:"not null" json:"address_id"`
	OrderCode      string           `gorm:"size:50;uniqueIndex;not null" json:"order_code"`
	Status         string           `gorm:"type:varchar(50);not null" json:"status"` // pending, confirmed, processing, shipped, delivered, cancelled, refunded
	Subtotal       float64          `gorm:"type:decimal(12,2);not null" json:"subtotal"`
	DiscountAmount float64          `gorm:"type:decimal(12,2);not null;default:0" json:"discount_amount"`
	ShippingFee    float64          `gorm:"type:decimal(12,2);not null" json:"shipping_fee"`
	TotalPrice     float64          `gorm:"type:decimal(12,2);not null" json:"total_price"`
	CouponID       *uint            `json:"coupon_id"`
	Note           string           `gorm:"type:text" json:"note"`
	PlacedAt       time.Time        `gorm:"not null;autoCreateTime" json:"placed_at"`
	UpdatedAt      time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	Items          []OrderItem      `gorm:"foreignKey:OrderID" json:"items"`
	StatusLogs     []OrderStatusLog `gorm:"foreignKey:OrderID" json:"status_logs"`
}

type OrderItem struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderID         uint           `gorm:"not null" json:"order_id"`
	ProductID       uint           `gorm:"not null" json:"product_id"`
	VariantID       *uint          `json:"variant_id"`
	Quantity        int            `gorm:"not null" json:"quantity"`
	UnitPrice       float64        `gorm:"type:decimal(12,2);not null" json:"unit_price"`
	ProductSnapshot datatypes.JSON `gorm:"not null" json:"product_snapshot"`
}

type OrderStatusLog struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `gorm:"not null" json:"order_id"`
	FromStatus    *string   `gorm:"type:varchar(50)" json:"from_status"`
	ToStatus      string    `gorm:"type:varchar(50);not null" json:"to_status"`
	ChangedByType string    `gorm:"type:varchar(20);not null" json:"changed_by_type"` // user, admin, system
	ChangedByID   uint      `gorm:"not null" json:"changed_by_id"`
	Note          string    `gorm:"type:text" json:"note"`
	ChangedAt     time.Time `gorm:"not null;autoCreateTime" json:"changed_at"`
}
