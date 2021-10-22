package product

type Service interface {
	GetProductsByCategoryID(CategoryID int) ([]Product, error)
}

type Repository interface {
	GetProductsByCategoryID(CategoryID int) ([]Product, error)
}
