package handler

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/usecase"
	"github.com/labstack/echo/v4"
)

type BankHandler struct {
	uc *usecase.BankUsecase
}

type WithdrawRequest struct {
	AccountID int `json:"account_id"`
	PlaceType int `json:"place_type"`
	Amount    int `json:"amount"`
}

func (h *BankHandler) Withdraw(c echo.Context) error {
	var req WithdrawRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	switch entity.PlaceType(req.PlaceType) {
	case entity.ConvenienceStore:
		if err := h.uc.WithdrawAtConvenienceStore(req.AccountID, req.AccountID); err != nil {
			return c.JSON(500, "error while drawing at convenience store")
		}
	default:
		return c.JSON(400, "place not found")
	}
	return c.JSON(200, "draw success")
}
