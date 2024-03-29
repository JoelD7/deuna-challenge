package models

import (
	"gorm.io/gorm"
	"time"
)

type PaymentStatus *string

var (
	PaymentStatusProcessing PaymentStatus = getStringPointer("processing")
	PaymentStatusFailed     PaymentStatus = getStringPointer("failed")
	PaymentStatusSuccess    PaymentStatus = getStringPointer("success")
)

type Payment struct {
	ID                string        `json:"id" gorm:"id;primaryKey"`
	MerchantAccountID *string       `json:"merchantAccountID" gorm:"merchant_account_id"`
	UserID            *string       `json:"userID" gorm:"user_id"`
	CardNumber        *int64        `json:"cardNumber" gorm:"card_number"`
	TransactionID     string        `json:"transactionID" gorm:"transaction_id"`
	Amount            *float64      `json:"amount" gorm:"amount"`
	Status            PaymentStatus `json:"status" gorm:"status"`
	FailureReason     string        `json:"failureReason" gorm:"failure_reason"`
	CreatedDate       *time.Time    `json:"createdDate" gorm:"created_date"`
	UpdatedDate       *time.Time    `json:"updatedDate" gorm:"updated_date"`
	Customer          User          `json:"-" gorm:"foreignKey:UserID"`
	Card              Card          `json:"-" gorm:"foreignKey:CardNumber"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = generateUUID()
	now := time.Now()
	p.CreatedDate = &now

	return
}
