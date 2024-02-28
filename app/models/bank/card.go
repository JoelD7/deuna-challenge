package bank

type Card struct {
	CardNumber int64   `json:"cardNumber"`
	CustomerID string  `json:"customerID"`
	AccountID  string  `json:"accountID"`
	Expiration string  `json:"expiration"`
	Type       string  `json:"type"`
	Vendor     string  `json:"vendor"`
	CCV        int     `json:"ccv"`
	Balance    float64 `json:"balance"`
}
