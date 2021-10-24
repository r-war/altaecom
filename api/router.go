package api

import (
	"AltaEcom/api/product"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	productController *product.Controller,
) {
	if productController == nil {
		panic("invalid parameter")
	}

	product := e.Group("/products")
	product.GET("", productController.GetProductsByCategoryID)

}
