package model

import "github.com/google/uuid"

type OrderItem struct {
	OrderId   uuid.UUID `db:"order_id" json:"order_id"`
	ProductID uuid.UUID `db:"product_id" json:"product_id"`
	Quantity  int32     `db:"quantity" json:"quantity"`
	Price     float64   `db:"price" json:"price"`
}
