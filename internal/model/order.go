package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId   uuid.UUID `db:"order_id" json:"order_id"`
	UserID    *uuid.UUID `db:"user_id" json:"user_id,omitempty"`
	StatusId  int32     `db:"status_id" json:"status_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
