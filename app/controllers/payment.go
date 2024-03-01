package controllers

import (
	"encoding/json"
	"github.com/JoelD7/deuna-challenge/app/db/repository"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/queue"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	sqliteClient = repository.NewSQLiteClient()
)

type processPaymentRequest struct {
	MerchantAccountID string `json:"merchantAccountID"`
}

type refundPaymentRequest struct {
	TransactionID string `json:"transactionID"`
}

func GetPaymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["paymentID"]

	getPayment := usecases.NewPaymentGetter(sqliteClient)

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

	createPayment := usecases.NewPaymentCreator(sqliteClient)
	getPayment := usecases.NewPaymentGetter(sqliteClient)

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

	if payment.MerchantAccountID == nil {
		return nil, models.ErrMissingMerchantAccountID
	}

	if payment.CardNumber == nil {
		return nil, models.ErrMissingCardNumber
	}

	return &payment, nil
}

func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if queue.IsEmpty() {
		w.WriteHeader(http.StatusNoContent)

		_, err := w.Write([]byte("No payments to process"))
		if err != nil {
			models.WriteErrorResponse(w, err)
		}

		return
	}

	var paymentReq processPaymentRequest

	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	if paymentReq.MerchantAccountID == "" {
		models.WriteErrorResponse(w, models.ErrMissingMerchantAccountID)
		return
	}

	processPayment := usecases.NewPaymentProcessor(sqliteClient)

	err = processPayment(r.Context(), paymentReq.MerchantAccountID)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RefundPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var refundReq refundPaymentRequest

	err := json.NewDecoder(r.Body).Decode(&refundReq)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	if refundReq.TransactionID == "" {
		models.WriteErrorResponse(w, models.ErrMissingTransactionID)
		return
	}

	refundPayment := usecases.NewPaymentRefunder(sqliteClient)

	err = refundPayment(r.Context(), refundReq.TransactionID)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
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
