package customer

import "AltaEcom/business/product"

type Service interface {
	getProductsByCategoryID(categoryId int) ([]product.Product, error)
}
