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

	jsonData, err := json.MarshalIndent(payment, "", "    ")
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
