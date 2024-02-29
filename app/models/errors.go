package models

import (
	"errors"
	"net/http"
)

type ErrResponse struct {
	Message string
	Status  int
}

var (
	ErrPaymentNotFound = errors.New("payment not found")

	statusByError = map[error]ErrResponse{
		ErrPaymentNotFound: {ErrPaymentNotFound.Error(), http.StatusNotFound},
	}
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "Internal server error"

	if res, ok := statusByError[err]; ok {
		status = res.Status
		message = res.Message
	}

	http.Error(w, message, status)
}
