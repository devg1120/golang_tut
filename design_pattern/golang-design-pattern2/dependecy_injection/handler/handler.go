package handler

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/usecase"
)

func NewBankHandler(uc *usecase.BankUsecase) *BankHandler {
	return &BankHandler{uc: uc}
}
