package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/bank/models"
)

type CardManager interface {
	GetCard(ctx context.Context, cardNumber int64) (*models.Card, error)
}

// NewCardValidator returns a function that checks if the passed in card is registered in the bank
func NewCardValidator(cm CardManager) func(ctx context.Context, card *models.Card) error {
	return func(ctx context.Context, card *models.Card) error {
		_, err := cm.GetCard(ctx, card.CardNumber)
		if err != nil {
			return err
		}

		return nil
	}
}
