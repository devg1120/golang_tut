package entity

import (
	"errors"
	"time"
)

type Bank struct {
	AccountID   string `gorm:"uniqueIndex"`
	Balance     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AccountType AccountType
}

type AccountType int

const (
	Regular AccountType = iota
	Premium
)

func (b *Bank) CanWithdraw(amount int) bool {
	return b.Balance < amount
}

// DomainServiceの方がいいかも
func (b *Bank) Withdraw(amount int) error {
	if !b.CanWithdraw(amount) {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Bank) ConvenienceStoreFee() int {
	fee := 50 // 基本手数料

	if b.AccountType == Premium {
		// PremiumPlumは手数料が半額
		fee /= 2
	}
	// 夜間営業時間は手数料が二倍
	if isNightTime() {
		fee *= 2
	}
	return fee
}

func isNightTime() bool {
	hour := time.Now().Hour()
	return hour < 6 || hour >= 22
}
