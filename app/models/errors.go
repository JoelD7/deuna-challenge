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
	ErrMissingMerchantAccountID     = errors.New("missing merchant account ID")
	ErrEqualCustomerIDAndMerchantID = errors.New("customer ID and merchant ID cannot be equal")
	ErrMissingCardNumber            = errors.New("missing card number")
	ErrInvalidCardNumber            = errors.New("invalid card number")
	ErrInvalidCardExpirationDate    = errors.New("invalid card expiration date")
	ErrInvalidCardExpirationFormat  = errors.New("invalid card expiration format. Make sure it's in the format MM/YY")
	ErrMissingCardExpiration        = errors.New("missing card expiry date")
	ErrMissingCCV                   = errors.New("missing CCV")
	ErrInvalidCCV                   = errors.New("invalid CCV. Must be a 3 digit number")
	ErrMissingCardType              = errors.New("missing card type")
	ErrInvalidCardType              = errors.New("invalid card type. Must be either 'debit' or 'credit'")
	ErrCardNotFound                 = errors.New("card not found")
	ErrInvalidCard                  = errors.New("invalid card")

	statusByError = map[error]ErrResponse{
		ErrPaymentNotFound:              {ErrPaymentNotFound.Error(), http.StatusNotFound},
		ErrInvalidAmount:                {ErrInvalidAmount.Error(), http.StatusBadRequest},
		ErrMissingAmount:                {ErrMissingAmount.Error(), http.StatusBadRequest},
		ErrMissingCustomerID:            {ErrMissingCustomerID.Error(), http.StatusBadRequest},
		ErrMissingMerchantAccountID:     {ErrMissingMerchantAccountID.Error(), http.StatusBadRequest},
		ErrEqualCustomerIDAndMerchantID: {ErrEqualCustomerIDAndMerchantID.Error(), http.StatusBadRequest},
		ErrMissingCardNumber:            {ErrMissingCardNumber.Error(), http.StatusBadRequest},
		ErrInvalidCardNumber:            {ErrInvalidCardNumber.Error(), http.StatusBadRequest},
		ErrInvalidCardExpirationDate:    {ErrInvalidCardExpirationDate.Error(), http.StatusBadRequest},
		ErrInvalidCardExpirationFormat:  {ErrInvalidCardExpirationFormat.Error(), http.StatusBadRequest},
		ErrMissingCardExpiration:        {ErrMissingCardExpiration.Error(), http.StatusBadRequest},
		ErrMissingCCV:                   {ErrMissingCCV.Error(), http.StatusBadRequest},
		ErrInvalidCCV:                   {ErrInvalidCCV.Error(), http.StatusBadRequest},
		ErrMissingCardType:              {ErrMissingCardType.Error(), http.StatusBadRequest},
		ErrInvalidCardType:              {ErrInvalidCardType.Error(), http.StatusBadRequest},
		ErrCardNotFound:                 {ErrCardNotFound.Error(), http.StatusNotFound},
		ErrInvalidCard:                  {ErrInvalidCard.Error(), http.StatusBadRequest},
	}
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "Internal server error"

	for mappedErr, responseErr := range statusByError {
		if errors.Is(err, mappedErr) {
			status = responseErr.Status
			message = err.Error()
		}
	}

	fmt.Println("Error: ", err.Error())

	http.Error(w, message, status)
}
