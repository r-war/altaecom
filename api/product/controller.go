package product

import (
	"AltaEcom/api/product/response"
	"AltaEcom/business/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetProductsByCategoryID(c echo.Context) error {

	query := c.QueryParam("category_id")
	id, _ := strconv.Atoi(query)
	products, err := controller.service.GetProductsByCategoryID(int(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"products": response.NewGetProductsByCategoryIDResponse(products).Products,
	})
}
