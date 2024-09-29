package repository

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
)

type Bank interface {
	Withdraw(accountID, amount int) error
	GetBank(accountID int) (*entity.Bank, error)
}
