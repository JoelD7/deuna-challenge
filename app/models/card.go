package models

type Card struct {
	CardNumber int64  `json:"cardNumber" gorm:"card_number;primaryKey"`
	CustomerID string `json:"customerID" gorm:"customer_id"`
	Expiration string `json:"expiration" gorm:"expiration"`
	Type       string `json:"type" gorm:"type"`
	Vendor     string `json:"vendor" gorm:"vendor"`
	CCV        int    `json:"ccv" gorm:"ccv"`
}
