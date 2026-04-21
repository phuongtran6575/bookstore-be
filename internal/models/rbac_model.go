package models

import (
	"time"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email        string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	FullName     string    `gorm:"size:300;not null" json:"full_name"`
	IsActive     bool      `gorm:"not null;default:true" json:"is_active"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Roles        []Role     `gorm:"many2many:admin_roles;" json:"roles"`
}

type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:150;uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Module      string `gorm:"size:100;not null" json:"module"` // catalog, order, user, payment, report, system
}

type AdminRole struct {
	AdminID    uint      `gorm:"primaryKey" json:"admin_id"`
	RoleID     uint      `gorm:"primaryKey" json:"role_id"`
	AssignedAt time.Time `gorm:"autoCreateTime" json:"assigned_at"`
	AssignedBy *uint     `json:"assigned_by"` // Admin ID who assigned this role
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
