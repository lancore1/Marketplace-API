package model

type StatusType string

type Status struct {
	StatusId int32      `db:"status_id" json:"status_id"`
	Status   StatusType `db:"status_type" json:"status_type"`
}

const (
	StatusPending   StatusType = "pending"
	StatusPaid      StatusType = "paid"
	StatusShipped   StatusType = "shipped"
	StatusDelivered StatusType = "delivered"
	StatusCancelled StatusType = "cancelled"
)
