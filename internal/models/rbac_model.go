package models

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AdminRole struct {
	AdminID int `json:"admin_id"`
	RoleID  int `json:"role_id"`
}

type RolePermission struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}
