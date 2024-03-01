package models

import (
	"time"
)

type Account struct {
	ID         string     `json:"id" gorm:"id"`
	CustomerID string     `json:"customerID" gorm:"customer_id"`
	Type       string     `json:"type" gorm:"type"`
	Balance    float64    `json:"balance" gorm:"balance"`
	OpenDate   *time.Time `json:"openDate" gorm:"open_date"`
	Customer   Customer   `json:"customer" gorm:"foreignKey:UserID"`
}
