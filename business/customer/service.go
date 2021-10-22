package customer

import "AltaEcom/business/product"

type service struct {
	productService product.Service
}

func newService(productService product.Service) Service {
	return &service{
		productService,
	}
}

func (s *service) GetProductsByCategoryID(categoryID int) ([]product.Product, error) {
		products, err := s.productService.GetProductsByCategoryID(categoryID)

		if err != nil {
			return []product.Product{}, nil
		}
		return products, err
}