package repository

import (
	"context"
	"errors"
	"github.com/JoelD7/deuna-challenge/bank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
)

type SQLiteClient struct{}

var (
	db *gorm.DB

	bankDB = os.Getenv("BANK_DB")
)

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open(bankDB), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func NewSQLiteClient() *SQLiteClient {
	return &SQLiteClient{}
}

func (cli *SQLiteClient) GetCard(ctx context.Context, cardNumber int64) (*models.Card, error) {
	var card models.Card

	err := db.Model(&models.Card{}).Preload(clause.Associations).First(&card, "card_number = ?", cardNumber).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrCardNotFound
	}

	if err != nil {
		return nil, err
	}

	return &card, nil
}
