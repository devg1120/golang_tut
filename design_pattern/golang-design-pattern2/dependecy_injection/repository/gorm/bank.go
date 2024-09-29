package gorm

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository"
	"gorm.io/gorm"
)

type bankImpl struct {
	db *gorm.DB
}

type GormParams struct {
	DB *gorm.DB
}

func NewBankImpl(p *GormParams) repository.Bank {
	return &bankImpl{db: p.DB}
}

func (b *bankImpl) Withdraw(accountID, amount int) error {
	var bank entity.Bank
	if err := transaction(b.db, func(tx *gorm.DB) error {
		if err := tx.Where("account_id = ?", accountID).First(&entity.Bank{}).Error; err != nil {
			return err
		}
		if err := bank.Withdraw(amount); err != nil {
			// logging
			return err
		}
		if err := tx.Model(&bank).Update("balance", bank.Balance).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		// logging
		return err
	}
	return nil
}

func (b *bankImpl) GetBank(accountID int) (*entity.Bank, error) {
	var bank entity.Bank
	if err := b.db.Where("account_id = ?", accountID).First(&bank).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}
