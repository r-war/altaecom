package order

import "time"

type Item struct {
	ID          int
	ProductID   int
	ProductName string
	Qty         int
	UpdateAt    time.Time
}

type Order struct {
	ID int
	Status bool
	CreatedAt time.Time
	UpdateAt time.Time
}

type OrderDetail struct {
	ID int
	Items []Item
	CreatedAt time.Time
}