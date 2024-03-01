package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/bank/models"
)

type TransactionManager interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error)
	ProcessTransaction(ctx context.Context, clientAccount models.Account, merchantAccountID string, amount float64) (string, error)
	GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error)
	GetAccount(ctx context.Context, accountID string) (*models.Account, error)
	UpdateAccount(ctx context.Context, account models.Account) error
}

type AccountManager interface {
	GetAccount(ctx context.Context, accountID string) (*models.Account, error)
}

func NewTransactionProcessor(tm TransactionManager, cm CardManager) func(ctx context.Context, cardNo int64, amount float64, merchantAccountID string) (string, error) {
	return func(ctx context.Context, cardNo int64, amount float64, merchantAccountID string) (string, error) {
		card, err := cm.GetCard(ctx, cardNo)
		if err != nil {
			return "", err
		}

		data, err := json.Marshal(card)

		if *card.Type == models.CardTypeDebit {
			fmt.Println("Card data: ", string(data))
			return tm.ProcessTransaction(ctx, card.Account, merchantAccountID, amount)
		}

		return "id", nil
	}
}

func NewTransactionGetter(tm TransactionManager) func(ctx context.Context, transactionID string) (*models.Transaction, error) {
	return func(ctx context.Context, transactionID string) (*models.Transaction, error) {
		return tm.GetTransaction(ctx, transactionID)
	}
}
