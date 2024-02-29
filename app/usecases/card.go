package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/app/models"
)

type CardManager interface {
	GetCard(ctx context.Context, cardNumber int64) (*models.Card, error)
	CreateCard(ctx context.Context, card models.Card) (int64, error)
}

func NewCardGetter(cm CardManager) func(ctx context.Context, cardNumber int64) (*models.Card, error) {
	return func(ctx context.Context, cardNumber int64) (*models.Card, error) {
		return cm.GetCard(ctx, cardNumber)
	}
}

func NewCardCreator(cm CardManager) func(ctx context.Context, card *models.Card) (int64, error) {
	return func(ctx context.Context, card *models.Card) (int64, error) {
		return cm.CreateCard(ctx, *card)
	}
}
