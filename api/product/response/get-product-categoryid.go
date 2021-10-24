package response

import "AltaEcom/business/product"

type getProductsByCategoryIDResponse struct {
	Products []GetProductsResponse `json:"products"`
}

func NewGetProductsByCategoryIDResponse(products []product.Product) getProductsByCategoryIDResponse {

	getProductsByCategoryIDResponse := getProductsByCategoryIDResponse{}

	for _, val := range products {
		var getProductsResponse GetProductsResponse

		getProductsResponse.ID = val.ID
		getProductsResponse.CategoryID = val.CategoryId
		getProductsResponse.Name = val.Name
		getProductsResponse.Price = val.Price
		getProductsResponse.Qty = val.Qty
		getProductsResponse.Description = val.Description
		getProductsResponse.Image = val.Image

		getProductsByCategoryIDResponse.Products = append(getProductsByCategoryIDResponse.Products, getProductsResponse)
	}

	if getProductsByCategoryIDResponse.Products == nil {
		getProductsByCategoryIDResponse.Products = []GetProductsResponse{}
	}

	return getProductsByCategoryIDResponse
}
