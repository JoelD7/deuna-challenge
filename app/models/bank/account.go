package bank

import (
	"time"
)

type Account struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customerID"`
	Type       string     `json:"type"`
	Balance    float64    `json:"balance"`
	OpenDate   *time.Time `json:"openDate"`
}
