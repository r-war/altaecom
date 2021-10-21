package product

type Service interface {
	getProductsByCategoryID(CategoryID int) ([]Product, error)

	findProductByID(id int) (*Product, error)

	insertProduct(product Product) error
}

type Repository interface {
	getProduct() (*Product, error)

	findProductByID(id int) (*Product, error)
}
