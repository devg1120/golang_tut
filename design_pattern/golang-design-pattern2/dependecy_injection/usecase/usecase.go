package usecase

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository"
)

func NewBankUsecase(repo repository.Bank) *BankUsecase {
	return &BankUsecase{repo: repo}
}
