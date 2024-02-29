package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/app/models"
)

type PaymentManager interface {
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	CreatePayment(ctx context.Context, payment models.Payment) (string, error)
}

func NewPaymentGetter(pm PaymentManager) func(ctx context.Context, paymentID string) (*models.Payment, error) {
	return func(ctx context.Context, paymentID string) (*models.Payment, error) {
		return pm.GetPayment(ctx, paymentID)
	}
}

func NewPaymentCreator(pm PaymentManager) func(ctx context.Context, payment *models.Payment) (string, error) {
	return func(ctx context.Context, payment *models.Payment) (string, error) {
		payment.Status = models.PaymentStatusProcessing

		return pm.CreatePayment(ctx, *payment)
	}
}
