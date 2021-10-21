package product

import "time"

type Product struct {
	ID          int
	CategoryId  int
	Name        string
	Price       int
	Qty         int
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
