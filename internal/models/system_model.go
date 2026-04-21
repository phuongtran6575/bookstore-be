package models

import (
	"time"

	"gorm.io/datatypes"
)

type SettingType string

const (
	SettingTypeString  SettingType = "string"
	SettingTypeNumber  SettingType = "number"
	SettingTypeBoolean SettingType = "boolean"
	SettingTypeJSON    SettingType = "json"
)

type AuditLog struct {
	ID         uint64         `gorm:"primaryKey" json:"id"`
	ActorType  ActorType      `gorm:"type:varchar(20);not null" json:"actor_type"` // user, admin, system
	ActorID    *uint          `json:"actor_id"`
	Action     string         `gorm:"size:200;not null" json:"action"`
	TargetType string         `gorm:"size:100;not null" json:"target_type"`
	TargetID   *uint          `json:"target_id"`
	Payload    datatypes.JSON `json:"payload"`
	IPAddress  string         `gorm:"size:45" json:"ip_address"`
	UserAgent  string         `gorm:"size:500" json:"user_agent"`
	CreatedAt  time.Time      `gorm:"not null;autoCreateTime" json:"created_at"`
}

type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Type      string         `gorm:"size:100;not null" json:"type"`
	Title     string         `gorm:"size:300;not null" json:"title"`
	Body      string         `gorm:"type:text;not null" json:"body"`
	Data      datatypes.JSON `json:"data"`
	IsRead    bool           `gorm:"not null;default:false" json:"is_read"`
	ReadAt    *time.Time     `json:"read_at"`
	SentAt    time.Time      `gorm:"not null;autoCreateTime" json:"sent_at"`
}

type Setting struct {
	Key         string      `gorm:"primaryKey;size:200" json:"key"`
	Value       string      `gorm:"type:text;not null" json:"value"`
	Type        SettingType `gorm:"type:varchar(20);not null" json:"type"` // string, number, boolean, json
	Description string      `gorm:"type:text" json:"description"`
	UpdatedBy   *uint       `json:"updated_by"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
