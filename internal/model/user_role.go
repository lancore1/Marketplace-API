package model

type RoleType string

type UserRole struct {
	RoleID int32    `db:"role_id" json:"role_id"`
	Role   RoleType `db:"role" json:"role_type"`
}

const (
	RoleAdmin RoleType = "admin"
	RoleUser  RoleType = "user"
)
