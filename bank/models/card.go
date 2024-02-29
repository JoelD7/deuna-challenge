package bank

type Card struct {
	CardNumber int64   `json:"cardNumber" gorm:"card_number"`
	CustomerID string  `json:"customerID" gorm:"customer_id"`
	AccountID  string  `json:"accountID" gorm:"account_id"`
	Expiration string  `json:"expiration" gorm:"expiration"`
	Type       string  `json:"type" gorm:"type"`
	Vendor     string  `json:"vendor" gorm:"vendor"`
	CCV        int     `json:"ccv" gorm:"ccv"`
	Balance    float64 `json:"balance" gorm:"balance"`
}