package controllers

import (
	"encoding/json"
	"github.com/JoelD7/deuna-challenge/app/db/repository"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPaymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["paymentID"]

	getPayment := usecases.NewPaymentGetter(repository.NewSQLiteClient())

	payment, err := getPayment(r.Context(), paymentID)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	writeJSONData(w, payment)
}

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	payment, err := validateCreatePaymentRequest(r)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	createPayment := usecases.NewPaymentCreator(repository.NewSQLiteClient())
	getPayment := usecases.NewPaymentGetter(repository.NewSQLiteClient())

	id, err := createPayment(r.Context(), payment)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	newPayment, err := getPayment(r.Context(), id)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJSONData(w, newPayment)
}

func validateCreatePaymentRequest(r *http.Request) (*models.Payment, error) {
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	if payment.Amount == nil {
		return nil, models.ErrMissingAmount
	}

	if payment.Amount != nil && *payment.Amount <= 0 {
		return nil, models.ErrInvalidAmount
	}

	if payment.CustomerID == nil {
		return nil, models.ErrMissingCustomerID
	}

	if payment.MerchantID == nil {
		return nil, models.ErrMissingMerchantID
	}

	if *payment.CustomerID == *payment.MerchantID {
		return nil, models.ErrEqualCustomerIDAndMerchantID
	}

	if payment.CardNumber == nil {
		return nil, models.ErrMissingCardNumber
	}

	return &payment, nil
}

func writeJSONData(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
