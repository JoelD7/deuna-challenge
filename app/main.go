package main

import (
	"github.com/JoelD7/deuna-challenge/app/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)

	r.HandleFunc("/payments/{paymentID}", controllers.GetPaymentHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/payments", controllers.CreatePaymentHandler).
		Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
