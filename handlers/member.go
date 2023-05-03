package handlers

import (
	"echo-test-1/models/webs"
	"echo-test-1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	appJSON = "application/json"
)

type MemberHandler interface {
	Members(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type MemberHandlerImpl struct {
	memberService services.MemberService
}

func NewMemberHandler(MemberService services.MemberService) MemberHandler {
	return &MemberHandlerImpl{memberService: MemberService}
}

func RespSuccess(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}

func (h *MemberHandlerImpl) Members(c echo.Context) error {
	members, err := h.memberService.GetAllMembers()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, members)
}

func (h *MemberHandlerImpl) Create(c echo.Context) error {
	m := new(webs.MemberDTO)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	newMember, err := h.memberService.CreateMember(*m)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return RespSuccess(c, http.StatusCreated, newMember)
}

func (h *MemberHandlerImpl) Update(c echo.Context) error {
	id := c.Param("id")
	parseMemberID, _ := strconv.ParseUint(id, 10, 32)

	m := new(webs.MemberDTO)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	updateMember, err := h.memberService.UpdateMember(uint(parseMemberID), *m)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return RespSuccess(c, http.StatusOK, updateMember)

}

func (h *MemberHandlerImpl) Delete(c echo.Context) error {
	id := c.Param("id")
	parseMemberID, _ := strconv.ParseUint(id, 10, 32)

	err := h.memberService.DeleteMember(uint(parseMemberID))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return RespSuccess(c, http.StatusOK, "Delete Has Been Successfull")
}
