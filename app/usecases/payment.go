package usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/queue"
	"io"
	"net/http"
	"os"
)

var bankHost = os.Getenv("BANK_HOST")

type PaymentManager interface {
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	CreatePayment(ctx context.Context, payment models.Payment) (string, error)
	UpdatePayment(ctx context.Context, payment models.Payment) error
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

		queue.Add(payment)

		return id, nil
	}
}

func NewPaymentProcessor(pm PaymentManager) func(ctx context.Context, merchantAccountID string) error {
	return func(ctx context.Context, merchantAccountID string) error {
		payment := queue.RemoveForMerchant(merchantAccountID)

		requestBody, err := json.Marshal(payment)
		if err != nil {
			return err
		}

		url := "http://" + bankHost + "/transaction"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			return fmt.Errorf("error sending request to bank: %w", err)
		}

		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			payment.Status = models.PaymentStatusFailed
			payment.FailureReason = string(responseBody)

			err = pm.UpdatePayment(ctx, *payment)
			if err != nil {
				return fmt.Errorf("error updating payment: %w", err)
			}

			return nil
		}

		payment.Status = models.PaymentStatusSuccess
		payment.TransactionID = string(responseBody)

		err = pm.UpdatePayment(ctx, *payment)
		if err != nil {
			return fmt.Errorf("error updating payment: %w", err)
		}

		return nil
	}
}

func NewPaymentRefunder(pm PaymentManager) func(ctx context.Context, transactionID string) error {
	return func(ctx context.Context, transactionID string) error {
		url := "http://" + bankHost + "/transaction/" + transactionID

		req, err := http.NewRequest(http.MethodPut, url, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request to bank: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error refunding payment")
		}

		return nil
	}
}
