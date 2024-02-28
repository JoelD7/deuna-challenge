package payment_platform

type Card struct {
	CardNumber int64  `json:"cardNumber"`
	CustomerID string `json:"customerID"`
	Expiration string `json:"expiration"`
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	CCV        int    `json:"ccv"`
}
