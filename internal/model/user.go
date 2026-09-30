package model

import "github.com/google/uuid"

type User struct {
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	Login        string    `db:"login" json:"login"`
	PasswordHash string    `db:"password_hash" json:"-"`
	FirstName    *string   `db:"first_name" json:"first_name,omitempty"`
	LastName     *string   `db:"last_name" json:"last_name,omitempty"`
	Email        string    `db:"email" json:"email"`
	RoleID       int32     `db:"role_id" json:"role_id"`
}
