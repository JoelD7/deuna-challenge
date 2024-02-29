package models

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string
	Status  int
}

var (
	ErrPaymentNotFound              = errors.New("payment not found")
	ErrInvalidAmount                = errors.New("invalid amount")
	ErrMissingAmount                = errors.New("missing amount")
	ErrMissingCustomerID            = errors.New("missing customer ID")
	ErrMissingMerchantID            = errors.New("missing merchant ID")
	ErrEqualCustomerIDAndMerchantID = errors.New("customer ID and merchant ID cannot be equal")
	ErrMissingCardNumber            = errors.New("missing card number")

	statusByError = map[error]ErrResponse{
		ErrPaymentNotFound:              {ErrPaymentNotFound.Error(), http.StatusNotFound},
		ErrInvalidAmount:                {ErrInvalidAmount.Error(), http.StatusBadRequest},
		ErrMissingAmount:                {ErrMissingAmount.Error(), http.StatusBadRequest},
		ErrMissingCustomerID:            {ErrMissingCustomerID.Error(), http.StatusBadRequest},
		ErrMissingMerchantID:            {ErrMissingMerchantID.Error(), http.StatusBadRequest},
		ErrEqualCustomerIDAndMerchantID: {ErrEqualCustomerIDAndMerchantID.Error(), http.StatusBadRequest},
		ErrMissingCardNumber:            {ErrMissingCardNumber.Error(), http.StatusBadRequest},
	}
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "Internal server error"

	if res, ok := statusByError[err]; ok {
		status = res.Status
		message = res.Message
	}

	fmt.Println("Error: ", err.Error())

	http.Error(w, message, status)
}
