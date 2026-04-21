package models

import (
	"time"

	"gorm.io/datatypes"
)

type PaymentMethodType string

const (
	PayTypeOffline    PaymentMethodType = "offline"
	PayTypeEWallet    PaymentMethodType = "e_wallet"
	PayTypeBank       PaymentMethodType = "bank"
	PayTypeCreditCard PaymentMethodType = "credit_card"
	PayTypeStoreCredit PaymentMethodType = "store_credit"
)

type PaymentStatus string

const (
	PaymentPending    PaymentStatus = "pending"
	PaymentProcessing PaymentStatus = "processing"
	PaymentPaid       PaymentStatus = "paid"
	PaymentFailed     PaymentStatus = "failed"
	PaymentRefunded   PaymentStatus = "refunded"
	PaymentExpired    PaymentStatus = "expired"
)

type RefundStatus string

const (
	RefundPending    RefundStatus = "pending"
	RefundProcessing RefundStatus = "processing"
	RefundCompleted  RefundStatus = "completed"
	RefundFailed     RefundStatus = "failed"
)

type RefundMethod string

const (
	RefundOriginalMethod RefundMethod = "original_method"
	RefundToWallet       RefundMethod = "store_wallet"
	RefundToBank         RefundMethod = "bank_transfer"
)

type TransactionType string

const (
	TxCredit TransactionType = "credit"
	TxDebit  TransactionType = "debit"
)

type WalletTransactionReason string

const (
	WalletReasonRefund      WalletTransactionReason = "refund"
	WalletReasonCashback    WalletTransactionReason = "cashback"
	WalletReasonPurchase    WalletTransactionReason = "purchase"
	WalletReasonAdminAdjust WalletTransactionReason = "admin_adjust"
	WalletReasonExpired     WalletTransactionReason = "expired"
)

type PaymentMethod struct {
	ID             uint              `gorm:"primaryKey" json:"id"`
	Code           string            `gorm:"size:50;uniqueIndex;not null" json:"code"` // cod, vnpay, momo, zalopay, bank_transfer, store_credit
	Name           string            `gorm:"size:200;not null" json:"name"`
	Type           PaymentMethodType `gorm:"type:varchar(50);not null" json:"type"` // offline, e_wallet, bank, credit_card, store_credit
	LogoURL        string            `gorm:"size:500" json:"logo_url"`
	Config         datatypes.JSON    `json:"config"` // encrypted in practice
	FeeType        FeeType           `gorm:"type:varchar(20);not null;default:none" json:"fee_type"` // none, percent, fixed
	FeeValue       float64           `gorm:"type:decimal(8,4);not null;default:0" json:"fee_value"`
	MinAmount      *float64       `gorm:"type:decimal(12,2)" json:"min_amount"`
	MaxAmount      *float64       `gorm:"type:decimal(12,2)" json:"max_amount"`
	IsActive       bool           `gorm:"not null;default:true" json:"is_active"`
	SortOrder      int            `json:"sort_order"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type Payment struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderID         uint           `gorm:"not null" json:"order_id"`
	PaymentMethodID uint           `gorm:"not null" json:"payment_method_id"`
	Status          PaymentStatus  `gorm:"type:varchar(50);not null" json:"status"` // pending, processing, paid, failed, refunded, expired
	Amount          float64        `gorm:"type:decimal(12,2);not null" json:"amount"`
	FeeAmount       float64        `gorm:"type:decimal(10,2);not null;default:0" json:"fee_amount"`
	TransactionRef  string         `gorm:"size:255" json:"transaction_ref"`
	GatewayRequest  datatypes.JSON `json:"gateway_request"`
	GatewayResponse datatypes.JSON `json:"gateway_response"`
	CallbackPayload datatypes.JSON `json:"callback_payload"`
	PaidAt          *time.Time     `json:"paid_at"`
	ExpiredAt       *time.Time     `json:"expired_at"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
}

type Refund struct {
	ID           uint         `gorm:"primaryKey" json:"id"`
	PaymentID    uint         `gorm:"not null" json:"payment_id"`
	OrderID      uint         `gorm:"not null" json:"order_id"`
	Amount       float64      `gorm:"type:decimal(12,2);not null" json:"amount"`
	Reason       string       `gorm:"type:text;not null" json:"reason"`
	Status       RefundStatus `gorm:"type:varchar(50);not null" json:"status"` // pending, processing, completed, failed
	RefundMethod RefundMethod `gorm:"type:varchar(50);not null" json:"refund_method"` // original_method, store_wallet, bank_transfer
	ProcessedBy  *uint        `json:"processed_by"` // Admin ID
	ProcessedAt  *time.Time   `json:"processed_at"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

type StoreWallet struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	Balance     float64   `gorm:"type:decimal(12,2);not null;default:0" json:"balance"`
	TotalEarned float64   `gorm:"type:decimal(14,2);not null;default:0" json:"total_earned"`
	TotalSpent  float64   `gorm:"type:decimal(14,2);not null;default:0" json:"total_spent"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type WalletTransaction struct {
	ID           uint                    `gorm:"primaryKey" json:"id"`
	WalletID     uint                    `gorm:"not null" json:"wallet_id"`
	Type         TransactionType         `gorm:"type:varchar(20);not null" json:"type"` // credit, debit
	Reason       WalletTransactionReason `gorm:"type:varchar(50);not null" json:"reason"` // refund, cashback, purchase, admin_adjust, expired
	Amount       float64                 `gorm:"type:decimal(12,2);not null" json:"amount"`
	BalanceAfter float64   `gorm:"type:decimal(12,2);not null" json:"balance_after"`
	RefOrderID   *uint     `json:"ref_order_id"`
	Note         string    `gorm:"size:500" json:"note"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
}
