package repository

import (
	"context"
	"github.com/JoelD7/deuna-challenge/app/models"
)

type PaymentRepository interface {
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
}
