package payment_platform

import (
	"time"
)

type Payment struct {
	ID            string     `json:"id"`
	MerchantID    string     `json:"merchantID"`
	CustomerID    string     `json:"customerID"`
	CardNumber    int        `json:"cardNumber"`
	TransactionID string     `json:"transactionID"`
	Amount        float64    `json:"amount"`
	Status        string     `json:"status"`
	FailureReason string     `json:"failureReason"`
	CreationDate  *time.Time `json:"creationDate"`
	UpdatedDate   *time.Time `json:"updatedDate"`
}
