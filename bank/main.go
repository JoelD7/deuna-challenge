package main

import (
	"fmt"
	"github.com/JoelD7/deuna-challenge/bank/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	bankHost = os.Getenv("BANK_HOST")
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)

	r.HandleFunc("/card", controllers.ValidateCardHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/card/{cardID}", controllers.GetCardHandler).
		Methods(http.MethodGet)

	r.HandleFunc("/transaction", controllers.CreateTransactionHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/transaction/{transactionID}", controllers.RefundTransactionHandler).
		Methods(http.MethodDelete)

	fmt.Println("Bank server running on", bankHost)
	log.Fatal(http.ListenAndServe(bankHost, r))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
