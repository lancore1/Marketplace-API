package model

type Category struct {
	CategoryID int32  `db:"category_id" json:"category_id,omitempty"`
	Title      string `db:"title" json:"title"`
}
