package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   uuid.UUID `db:"product_id" json:"product_id"`
	Title       string    `db:"title" json:"title"`
	CategoryID  *int32    `db:"category_id" json:"category_id,omitempty"`
	Description *string   `db:"description" json:"description,omitempty"`
	Price       float64   `db:"price" json:"price"`
	Quantity    int32     `db:"quantity" json:"quantity"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
