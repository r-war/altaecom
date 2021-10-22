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

func NewProduct(
	ID          int,
	CategoryId  int,
	Name        string,
	Price       int,
	Qty         int,
	Description string,
	Image       string,
	CreatedAt   time.Time,
	UpdatedAt   time.Time,
	DeletedAt   time.Time,) Product {

		return Product{
			ID         : ID,
			CategoryId : CategoryId,
			Name       : Name,
			Price      : Price,
			Qty        : Qty,
			Description: Description,
			Image      : Image,
			CreatedAt  : CreatedAt,
			UpdatedAt  : UpdatedAt,
			DeletedAt  : DeletedAt,
	}
}