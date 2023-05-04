package handlers

import (
	"echo-test-1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	dislike = "DISLIKE"
	like    = "LIKE"
)

type LikeReviewHandler interface {
	Reaction(c echo.Context) error
}

type LikeReviewHandlerImpl struct {
	likeReviewService services.LikeReviewService
}

func NewLikeReviewHandler(LikeReviewService services.LikeReviewService) LikeReviewHandler {
	return &LikeReviewHandlerImpl{likeReviewService: LikeReviewService}
}

func (h *LikeReviewHandlerImpl) Reaction(c echo.Context) error {
	reviewID := c.Param("review_id")
	parseReviewID, _ := strconv.ParseUint(reviewID, 10, 32)
	memberID := c.Request().Header.Get("member_id")
	parseMemberID, _ := strconv.ParseUint(memberID, 10, 32)
	likeQuery := c.QueryParam("reaction")

	if likeQuery == dislike {
		err := h.likeReviewService.DislikeReview(uint(parseReviewID))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusAccepted, "Delete Has Been Successfull")
	}
	newReaction, err := h.likeReviewService.LikeReview(uint(parseReviewID), uint(parseMemberID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return RespSuccess(c, http.StatusCreated, newReaction)
}
