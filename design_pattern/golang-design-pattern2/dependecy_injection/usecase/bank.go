package usecase

import (
	"errors"
	"fmt"

	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository"
)

type BankUsecase struct {
	repo repository.Bank
}

func (u *BankUsecase) WithdrawAtConvenienceStore(accountID int, amount int) error {
	bank, err := u.repo.GetBank(accountID)
	if err != nil {
		return fmt.Errorf("failed to get bank account %d: %v", accountID, err)
	}

	// 手数料の計算
	fee := bank.ConvenienceStoreFee()

	totalAmount := fee + amount

	if !bank.CanWithdraw(totalAmount) {
		return errors.New("insufficient funds for withdrawal and fee")
	}

	if err := u.repo.Withdraw(accountID, totalAmount); err != nil {
		return fmt.Errorf("failed to withdraw from account %d: %v", accountID, err)
	}

	return nil
}
