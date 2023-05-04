package handlers

import (
	"echo-test-1/models/webs"
	"echo-test-1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	Products(c echo.Context) error
	Create(c echo.Context) error
	Detail(c echo.Context) error
}

type ProductHandlerImpl struct {
	productService services.ProductService
}

func NewProductHandler(ProductService services.ProductService) ProductHandler {
	return &ProductHandlerImpl{productService: ProductService}
}

func (h *ProductHandlerImpl) Products(c echo.Context) error {
	products, err := h.productService.GetAllProducts()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return RespSuccess(c, http.StatusOK, products)
}

func (h *ProductHandlerImpl) Create(c echo.Context) error {
	p := new(webs.ProductDTO)

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	newProduct, err := h.productService.CreateProduct(*p)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return RespSuccess(c, http.StatusCreated, newProduct)
}

func (h *ProductHandlerImpl) Detail(c echo.Context) error {
	productID := c.Param("product_id")
	parseProductID, _ := strconv.ParseUint(productID, 10, 32)
	product, err := h.productService.DetailProduct(uint(parseProductID))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, product)
}
