package model

import "github.com/google/uuid"

type CartItem struct {
	CartId    uuid.UUID `db:"cart_id" json:"cart_id"`
	ProductID uuid.UUID `db:"product_id" json:"product_id"`
	Quantity  int32     `db:"quantity" json:"quantity"`
}
