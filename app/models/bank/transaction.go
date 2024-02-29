package bank

type Transaction struct {
	ID                 string  `json:"id" gorm:"id"`
	AccountID          string  `json:"accountID" gorm:"account_id"`
	RecipientAccountID string  `json:"recipientAccountID" gorm:"recipient_account_id"`
	Amount             float64 `json:"amount" gorm:"amount"`
	Type               string  `json:"type" gorm:"type"`
	Date               string  `json:"date" gorm:"date"`
}
