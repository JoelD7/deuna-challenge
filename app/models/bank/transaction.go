package bank

type Transaction struct {
	ID                 string  `json:"id"`
	AccountID          string  `json:"accountID"`
	RecipientAccountID string  `json:"recipientAccountID"`
	Amount             float64 `json:"amount"`
	Type               string  `json:"type"`
	Date               string  `json:"date"`
}
