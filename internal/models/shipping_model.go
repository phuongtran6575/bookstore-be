package models

import (
	"time"

	"gorm.io/datatypes"
)

type ShippingMethodType string

const (
	ShipTypeStandard     ShippingMethodType = "standard"
	ShipTypeExpress      ShippingMethodType = "express"
	ShipTypeSameDay      ShippingMethodType = "same_day"
	ShipTypePickupInStore ShippingMethodType = "pickup_in_store"
)

type ShipmentStatus string

const (
	ShipmentPending         ShipmentStatus = "pending"
	ShipmentPickedUp        ShipmentStatus = "picked_up"
	ShipmentInTransit       ShipmentStatus = "in_transit"
	ShipmentOutForDelivery  ShipmentStatus = "out_for_delivery"
	ShipmentDelivered       ShipmentStatus = "delivered"
	ShipmentFailed          ShipmentStatus = "failed"
	ShipmentReturned        ShipmentStatus = "returned"
)

type Carrier struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	Code                string         `gorm:"size:50;uniqueIndex;not null" json:"code"` // ghn, ghtk, vtp, vnpost, grab_express
	Name                string         `gorm:"size:200;not null" json:"name"`
	LogoURL             string         `gorm:"size:500" json:"logo_url"`
	APIEndpoint         string         `gorm:"size:500" json:"api_endpoint"`
	APIConfig           datatypes.JSON `json:"api_config"`
	TrackingURLTemplate string         `gorm:"size:500" json:"tracking_url_template"`
	WebhookSecret       string         `gorm:"size:255" json:"-"`
	IsActive            bool           `gorm:"not null;default:true" json:"is_active"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type ShippingMethod struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CarrierID    uint      `gorm:"not null" json:"carrier_id"`
	Carrier      Carrier   `gorm:"foreignKey:CarrierID" json:"carrier"`
	Code         string             `gorm:"size:100;uniqueIndex;not null" json:"code"`
	Name         string             `gorm:"size:200;not null" json:"name"`
	Type         ShippingMethodType `gorm:"type:varchar(50);not null" json:"type"` // standard, express, same_day, pickup_in_store
	MinDays      *int               `json:"min_days"`
	MaxDays      *int      `json:"max_days"`
	MaxWeightKG  *float64  `gorm:"type:decimal(6,2)" json:"max_weight_kg"`
	IsActive     bool      `gorm:"not null;default:true" json:"is_active"`
	SortOrder    int       `json:"sort_order"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ShippingZone struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:200;not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Provinces   datatypes.JSON `json:"provinces"`
	Districts   datatypes.JSON `json:"districts"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type ShippingRate struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ShippingMethodID uint      `gorm:"not null" json:"shipping_method_id"`
	ZoneID           uint      `gorm:"not null" json:"zone_id"`
	MinWeightG       int       `gorm:"not null" json:"min_weight_g"`
	MaxWeightG       int       `gorm:"not null" json:"max_weight_g"`
	BaseFee          float64   `gorm:"type:decimal(10,2);not null" json:"base_fee"`
	Per500gFee       float64   `gorm:"type:decimal(10,2);not null;default:0" json:"per_500g_fee"`
	FreeShippingFrom *float64  `gorm:"type:decimal(12,2)" json:"free_shipping_from"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Shipment struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	OrderID             uint           `gorm:"not null" json:"order_id"`
	ShippingMethodID    uint           `gorm:"not null" json:"shipping_method_id"`
	CarrierID           uint           `gorm:"not null" json:"carrier_id"`
	TrackingCode        string         `gorm:"size:200;uniqueIndex" json:"tracking_code"`
	ShippingFeeCharged  float64        `gorm:"type:decimal(10,2);not null" json:"shipping_fee_charged"`
	WeightG             *int           `json:"weight_g"`
	Status              ShipmentStatus `gorm:"type:varchar(50);not null" json:"status"` // pending, picked_up, in_transit, out_for_delivery, delivered, failed, returned
	CarrierStatusRaw    datatypes.JSON `json:"carrier_status_raw"`
	ShippedAt           *time.Time     `json:"shipped_at"`
	DeliveredAt         *time.Time     `json:"delivered_at"`
	FailedReason        string         `gorm:"type:text" json:"failed_reason"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type ShipmentTrackingEvent struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ShipmentID  uint           `gorm:"not null" json:"shipment_id"`
	StatusCode  string         `gorm:"size:100;not null" json:"status_code"`
	Description string         `gorm:"type:text" json:"description"`
	Location    string         `gorm:"size:300" json:"location"`
	OccurredAt  time.Time      `gorm:"not null" json:"occurred_at"`
	RawPayload  datatypes.JSON `json:"raw_payload"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
}
