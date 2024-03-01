package models

import (
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeWithdraw TransactionType = "withdraw"
	TransactionTypeTransfer TransactionType = "transfer"
	TransactionTypeCredit   TransactionType = "credit_card_payment"

	TransactionStatusSuccess  TransactionStatus = "success"
	TransactionStatusRefunded TransactionStatus = "refunded"
)

type Transaction struct {
	ID                 string            `json:"id" gorm:"id"`
	AccountID          string            `json:"accountID" gorm:"account_id"`
	RecipientAccountID string            `json:"recipientAccountID" gorm:"recipient_account_id"`
	CreditCardNumber   int64             `json:"creditCardNumber" gorm:"credit_card_number"`
	Amount             float64           `json:"amount" gorm:"amount"`
	Type               TransactionType   `json:"type" gorm:"type"`
	Status             TransactionStatus `json:"status" gorm:"status"`
	CreatedDate        *time.Time        `json:"createdDate" gorm:"created_date"`
	Account            Account           `json:"account" gorm:"foreignKey:AccountID"`
	RecipientAccount   Account           `json:"recipientAccount" gorm:"foreignKey:RecipientAccountID"`
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
