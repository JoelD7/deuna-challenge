package repository

import (
	"context"
	"errors"
	"github.com/JoelD7/deuna-challenge/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SQLiteClient struct{}

var (
	db *gorm.DB
)

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("deuna-db.sqlt"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func NewSQLiteClient() *SQLiteClient {
	return &SQLiteClient{}
}

func (cli *SQLiteClient) GetPayment(ctx context.Context, paymentID string) (*models.Payment, error) {
	var payment models.Payment

	err := db.Model(&models.Payment{}).Preload(clause.Associations).First(&payment, "id = ?", paymentID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrPaymentNotFound
	}

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (cli *SQLiteClient) CreatePayment(ctx context.Context, payment models.Payment) (string, error) {
	result := db.Create(&payment)

	if result.Error != nil {
		return "", result.Error
	}

	if payment.ID == "" {
		return "", errors.New("payment ID is nil")
	}

	return payment.ID, nil
}
