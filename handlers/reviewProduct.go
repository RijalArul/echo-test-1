package handlers

import (
	"echo-test-1/models/webs"
	"echo-test-1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewProductHandler interface {
	Create(c echo.Context) error
	Reviews(c echo.Context) error
}

type ReviewProductHandlerImpl struct {
	reviewProductService services.ReviewProductService
}

func NewReviewProductHandler(ReviewProductService services.ReviewProductService) ReviewProductHandler {
	return &ReviewProductHandlerImpl{reviewProductService: ReviewProductService}
}

func (h *ReviewProductHandlerImpl) Create(c echo.Context) error {
	r := new(webs.ReviewProductDTO)
	productID := c.Param("product_id")
	parseProductID, _ := strconv.ParseUint(productID, 10, 32)

	memberID := c.Request().Header.Get("member_id")
	parseMemberID, _ := strconv.ParseUint(memberID, 10, 32)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	newReview, err := h.reviewProductService.CreateReview(*r, uint(parseMemberID), uint(parseProductID))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return RespSuccess(c, http.StatusCreated, newReview)
}

func (h *ReviewProductHandlerImpl) Reviews(c echo.Context) error {
	productID := c.Param("product_id")
	parseProductID, _ := strconv.ParseUint(productID, 10, 32)

	reviewsProduct, err := h.reviewProductService.ReviewsProduct(uint(parseProductID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return RespSuccess(c, http.StatusOK, reviewsProduct)
}
