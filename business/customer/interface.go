package customer

import "AltaEcom/business/product"

type Service interface {
	GetProductsByCategoryID(categoryId int) ([]product.Product, error)
}
