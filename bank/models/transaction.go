package models

import (
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type TransactionType string

const (
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeWithdraw TransactionType = "withdraw"
	TransactionTypeTransfer TransactionType = "transfer"
)

type Transaction struct {
	ID                 string          `json:"id" gorm:"id"`
	AccountID          string          `json:"accountID" gorm:"account_id"`
	RecipientAccountID string          `json:"recipientAccountID" gorm:"recipient_account_id"`
	Amount             float64         `json:"amount" gorm:"amount"`
	Type               TransactionType `json:"type" gorm:"type"`
	CreatedDate        *time.Time      `json:"createdDate" gorm:"created_date"`
	Account            Account         `json:"account" gorm:"foreignKey:AccountID"`
	RecipientAccount   Account         `json:"recipientAccount" gorm:"foreignKey:RecipientAccountID"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = generateUUID()
	now := time.Now()
	t.CreatedDate = &now

	return
}

func generateUUID() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 15)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
