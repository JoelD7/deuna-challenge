package models

type CardType *string

var (
	Debit  CardType = getStringPointer("debit")
	Credit CardType = getStringPointer("credit")
)

type Card struct {
	CardNumber int64    `json:"cardNumber" gorm:"card_number"`
	CustomerID *string  `json:"customerID" gorm:"customer_id"`
	AccountID  *string  `json:"accountID" gorm:"account_id"`
	Expiration *string  `json:"expiration" gorm:"expiration"`
	Type       CardType `json:"type" gorm:"type"`
	Vendor     *string  `json:"vendor" gorm:"vendor"`
	CCV        int      `json:"ccv" gorm:"ccv"`
	Balance    float64  `json:"balance" gorm:"balance"`
	Account    Account  `json:"account" gorm:"foreignKey:AccountID"`
}

func getStringPointer(s string) *string {
	return &s
}
