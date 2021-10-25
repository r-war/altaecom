package product

import (
	"AltaEcom/api/common"
	"AltaEcom/api/product/request"
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
		"products": response.NewGetProductsResponse(products).Products,
	})
}

func (c *Controller) GetProducts(ctx echo.Context) error {
	products, err := c.service.GetProducts()
	if err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}
	return ctx.JSON(common.BusinessResponse(
		response.NewGetProductsResponse(products).Products,
	))
}

func (c *Controller) FindProductByid(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.service.FindProductByid(id)
	if err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}
	res := response.NewGetProductResponse(*product)

	return ctx.JSON(common.NewSuccessResponse(res))

}

func (c *Controller) InsertProduct(ctx echo.Context) error {
	var err error
	insertProduct := new(request.ProductRequest)
	if err = ctx.Bind(insertProduct); err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}
	if err = c.service.InsertProduct(insertProduct); err != nil {
		return ctx.JSON(common.NewErrorBusinessResponse(err))
	}

	return ctx.JSON(common.NewSuccessResponseWithoutData())
}
