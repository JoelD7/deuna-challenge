package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/queue"
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

		id, err := pm.CreatePayment(ctx, *payment)
		if err != nil {
			return "", err
		}

		payment.ID = id

		queue.Enqueue(payment)

		return id, nil
	}
}
