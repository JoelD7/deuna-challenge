package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/bank/models"
)

type TransactionManager interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error)
	ProcessDebitTransaction(ctx context.Context, clientAccount models.Account, merchantAccountID string, amount float64) (string, error)
	ProcessCreditTransaction(ctx context.Context, clientCard *models.Card, merchantAccountID string, amount float64) (string, error)
	RefundDebitTransaction(ctx context.Context, transaction *models.Transaction) error
	RefundCreditTransaction(ctx context.Context, transaction *models.Transaction) error
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

		if *card.Type == models.CardTypeDebit {
			return tm.ProcessDebitTransaction(ctx, card.Account, merchantAccountID, amount)
		}

		return tm.ProcessCreditTransaction(ctx, card, merchantAccountID, amount)
	}
}

func NewTransactionRefunder(tm TransactionManager) func(ctx context.Context, transactionID string) error {
	return func(ctx context.Context, transactionID string) error {
		transaction, err := tm.GetTransaction(ctx, transactionID)
		if err != nil {
			return err
		}

		if transaction.Type == models.TransactionTypeTransfer {
			return tm.RefundDebitTransaction(ctx, transaction)
		}

		return tm.RefundCreditTransaction(ctx, transaction)
	}
}

func NewTransactionGetter(tm TransactionManager) func(ctx context.Context, transactionID string) (*models.Transaction, error) {
	return func(ctx context.Context, transactionID string) (*models.Transaction, error) {
		return tm.GetTransaction(ctx, transactionID)
	}
}
