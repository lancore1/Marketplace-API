package model

import "github.com/google/uuid"

type Cart struct {
	CartId uuid.UUID `db:"cart_id" json:"cart_id"`
	UserID uuid.UUID `db:"user_id" json:"user_id"`
}
