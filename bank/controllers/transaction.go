package controllers

import (
	"encoding/json"
	"github.com/JoelD7/deuna-challenge/bank/db/repository"
	"github.com/JoelD7/deuna-challenge/bank/models"
	"github.com/JoelD7/deuna-challenge/bank/usecases"
	"net/http"
)

type transactionRequest struct {
	CardNumber        int64   `json:"cardNumber"`
	Amount            float64 `json:"amount"`
	MerchantAccountID string  `json:"merchantAccountID"`
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tr transactionRequest

	err := json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	err = validateTransactionFields(&tr)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	processTransaction := usecases.NewTransactionProcessor(repository.NewSQLiteClient(), repository.NewSQLiteClient())

	id, err := processTransaction(r.Context(), tr.CardNumber, tr.Amount, tr.MerchantAccountID)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(id))
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}
}

func validateTransactionFields(tr *transactionRequest) error {
	if tr.CardNumber == 0 {
		return models.ErrMissingCardNumber
	}

	if tr.Amount == 0 {
		return models.ErrMissingAmount
	}

	if tr.MerchantAccountID == "" {
		return models.ErrMissingMerchantAccountID
	}

	return nil
}
