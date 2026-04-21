package models

import (
	"time"

	"gorm.io/datatypes"
)

type LoyaltyAccount struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	Tier           string    `gorm:"type:varchar(20);not null" json:"tier"` // bronze, silver, gold, platinum
	TotalPoints    int       `gorm:"not null;default:0" json:"total_points"`
	LifetimePoints int       `gorm:"not null;default:0" json:"lifetime_points"`
	LifetimeSpend  float64   `gorm:"type:decimal(14,2);not null;default:0" json:"lifetime_spend"`
	TierExpiresAt  *time.Time `json:"tier_expires_at"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type LoyaltyTransaction struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	LoyaltyAccountID uint      `gorm:"not null" json:"loyalty_account_id"`
	Type             string    `gorm:"type:varchar(50);not null" json:"type"` // earn, redeem, expire, admin_adjust
	Points           int       `gorm:"not null" json:"points"`
	BalanceAfter     int       `gorm:"not null" json:"balance_after"`
	RefOrderID       *uint     `json:"ref_order_id"`
	ExpiresAt        *time.Time `json:"expires_at"`
	Note             string    `gorm:"size:500" json:"note"`
	CreatedAt        time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}

type StoreLocation struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Code              string         `gorm:"size:20;uniqueIndex;not null" json:"code"`
	Name              string         `gorm:"size:300;not null" json:"name"`
	Type              string         `gorm:"type:varchar(20);not null" json:"type"` // store, warehouse
	Address           string         `gorm:"size:500;not null" json:"address"`
	Ward              string         `gorm:"size:200" json:"ward"`
	District          string         `gorm:"size:200;not null" json:"district"`
	City              string         `gorm:"size:200;not null" json:"city"`
	Lat               *float64       `gorm:"type:decimal(10,7)" json:"lat"`
	Lng               *float64       `gorm:"type:decimal(10,7)" json:"lng"`
	Phone             string         `gorm:"size:20" json:"phone"`
	OpeningHours      datatypes.JSON `json:"opening_hours"`
	IsPickupAvailable bool           `gorm:"not null;default:true" json:"is_pickup_available"`
	IsActive          bool           `gorm:"not null;default:true" json:"is_active"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type LocationInventory struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	LocationID     uint      `gorm:"not null;uniqueIndex:idx_location_variant" json:"location_id"`
	VariantID      uint      `gorm:"not null;uniqueIndex:idx_location_variant" json:"variant_id"`
	StockOnHand    int       `gorm:"not null;default:0" json:"stock_on_hand"`
	StockReserved  int       `gorm:"not null;default:0" json:"stock_reserved"`
	StockAvailable int       `gorm:"not null;default:0" json:"stock_available"` // Note: In GORM we might need to handle the GENERATED aspect manually or via DB triggers
	ReorderPoint   *int      `json:"reorder_point"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type InventoryLog struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	LocationInventoryID uint      `gorm:"not null" json:"location_inventory_id"`
	MovementType        string    `gorm:"type:varchar(50);not null" json:"movement_type"` // purchase_in, sale_out, etc.
	QuantityChange      int       `gorm:"not null" json:"quantity_change"`
	StockBefore         int       `gorm:"not null" json:"stock_before"`
	StockAfter          int       `gorm:"not null" json:"stock_after"`
	RefType             string    `gorm:"size:50" json:"ref_type"`
	RefID               *uint     `json:"ref_id"`
	ActorType           string    `gorm:"type:varchar(20)" json:"actor_type"` // user, admin, system
	ActorID             *uint     `json:"actor_id"`
	Note                string    `gorm:"type:text" json:"note"`
	CreatedAt           time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
