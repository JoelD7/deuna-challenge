package repository

import (
	"context"
	"errors"
	"github.com/JoelD7/deuna-challenge/bank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"time"
)

type SQLiteClient struct {
	conn *gorm.DB
}

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
	return &SQLiteClient{
		conn: db,
	}
}

func (cli *SQLiteClient) SetConnection(conn *gorm.DB) {
	cli.conn = conn
}

func (cli *SQLiteClient) GetCard(ctx context.Context, cardNumber int64) (*models.Card, error) {
	var card models.Card

	err := cli.conn.Model(&models.Card{}).Preload(clause.Associations).First(&card, "card_number = ?", cardNumber).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrCardNotFound
	}

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (cli *SQLiteClient) UpdateCard(ctx context.Context, card models.Card) error {
	return cli.conn.Save(&card).Error
}

func (cli *SQLiteClient) UpdateAccount(ctx context.Context, account models.Account) error {
	return cli.conn.Save(&account).Error
}

func (cli *SQLiteClient) GetAccount(ctx context.Context, accountID string) (*models.Account, error) {
	var account models.Account

	err := cli.conn.Model(&models.Account{}).Preload(clause.Associations).First(&account, "id = ?", accountID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (cli *SQLiteClient) ProcessDebitTransaction(ctx context.Context, clientAccount models.Account, merchantAccountID string, amount float64) (string, error) {
	transactionID := ""

	err := cli.conn.Transaction(func(tx *gorm.DB) error {
		cli.SetConnection(tx)

		merchantAccount, err := cli.GetAccount(ctx, merchantAccountID)
		if err != nil {
			return err
		}

		now := time.Now()

		newTransaction := &models.Transaction{
			AccountID:          clientAccount.ID,
			RecipientAccountID: merchantAccount.ID,
			Amount:             amount,
			Type:               models.TransactionTypeTransfer,
			CreatedDate:        &now,
		}

		id, err := cli.CreateTransaction(ctx, *newTransaction)
		if err != nil {
			return err
		}

		clientAccount.Balance -= amount
		if err = cli.UpdateAccount(ctx, clientAccount); err != nil {
			return err
		}

		merchantAccount.Balance += amount
		if err = cli.UpdateAccount(ctx, *merchantAccount); err != nil {
			return err
		}

		transactionID = id

		return nil
	})

	return transactionID, err
}

func (cli *SQLiteClient) ProcessCreditTransaction(ctx context.Context, clientCard *models.Card, merchantAccountID string, amount float64) (string, error) {
	transactionID := ""

	err := cli.conn.Transaction(func(tx *gorm.DB) error {
		cli.SetConnection(tx)

		merchantAccount, err := cli.GetAccount(ctx, merchantAccountID)
		if err != nil {
			return err
		}

		now := time.Now()

		newTransaction := &models.Transaction{
			CreditCardNumber:   clientCard.CardNumber,
			RecipientAccountID: merchantAccount.ID,
			Amount:             amount,
			Type:               models.TransactionTypeCredit,
			CreatedDate:        &now,
		}

		id, err := cli.CreateTransaction(ctx, *newTransaction)
		if err != nil {
			return err
		}

		clientCard.Balance -= amount
		if err = cli.UpdateCard(ctx, *clientCard); err != nil {
			return err
		}

		merchantAccount.Balance += amount
		if err = cli.UpdateAccount(ctx, *merchantAccount); err != nil {
			return err
		}

		transactionID = id

		return nil
	})

	return transactionID, err
}

func (cli *SQLiteClient) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	err := cli.conn.Create(&transaction).Error

	if err != nil {
		return "", err
	}

	return transaction.ID, nil
}

func (cli *SQLiteClient) GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error) {
	var transaction models.Transaction

	err := cli.conn.Model(&models.Transaction{}).Preload(clause.Associations).First(&transaction, "id = ?", transactionID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, models.ErrTransactionNotFound
	}

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
