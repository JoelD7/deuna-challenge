package payment_platform

import (
	"time"
)

type Payment struct {
	ID            string     `json:"id" gorm:"id"`
	MerchantID    string     `json:"merchantID" gorm:"merchant_id"`
	CustomerID    string     `json:"customerID" gorm:"customer_id"`
	CardNumber    int        `json:"cardNumber" gorm:"card_number"`
	TransactionID string     `json:"transactionID" gorm:"transaction_id"`
	Amount        float64    `json:"amount" gorm:"amount"`
	Status        string     `json:"status" gorm:"status"`
	FailureReason string     `json:"failureReason" gorm:"failure_reason"`
	CreatedDate   *time.Time `json:"createdDate" gorm:"created_date"`
	UpdatedDate   *time.Time `json:"updatedDate" gorm:"updated_date"`
}
