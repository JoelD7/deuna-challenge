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
	ErrMissingCardNumber           = errors.New("missing card number")
	ErrInvalidCardNumber           = errors.New("invalid card number")
	ErrInvalidCardExpirationDate   = errors.New("invalid card expiration date")
	ErrInvalidCardExpirationFormat = errors.New("invalid card expiration format. Make sure it's in the format MM/YY")
	ErrMissingCardExpiration       = errors.New("missing card expiry date")
	ErrMissingCCV                  = errors.New("missing CCV")
	ErrInvalidCCV                  = errors.New("invalid CCV. Must be a 3 digit number")
	ErrMissingCardType             = errors.New("missing card type")
	ErrInvalidCardType             = errors.New("invalid card type. Must be either 'debit' or 'credit'")
	ErrCardNotFound                = errors.New("card not found")

	statusByError = map[error]ErrResponse{
		ErrMissingCardNumber:           {ErrMissingCardNumber.Error(), http.StatusBadRequest},
		ErrInvalidCardNumber:           {ErrInvalidCardNumber.Error(), http.StatusBadRequest},
		ErrInvalidCardExpirationDate:   {ErrInvalidCardExpirationDate.Error(), http.StatusBadRequest},
		ErrInvalidCardExpirationFormat: {ErrInvalidCardExpirationFormat.Error(), http.StatusBadRequest},
		ErrMissingCardExpiration:       {ErrMissingCardExpiration.Error(), http.StatusBadRequest},
		ErrMissingCCV:                  {ErrMissingCCV.Error(), http.StatusBadRequest},
		ErrInvalidCCV:                  {ErrInvalidCCV.Error(), http.StatusBadRequest},
		ErrMissingCardType:             {ErrMissingCardType.Error(), http.StatusBadRequest},
		ErrInvalidCardType:             {ErrInvalidCardType.Error(), http.StatusBadRequest},
		ErrCardNotFound:                {ErrCardNotFound.Error(), http.StatusNotFound},
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
