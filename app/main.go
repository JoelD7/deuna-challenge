package main

import (
	"github.com/JoelD7/deuna-challenge/app/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	appURL = os.Getenv("APP_URL")
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)

	r.HandleFunc("/payments/{paymentID}", controllers.GetPaymentHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/payments", controllers.CreatePaymentHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/card", controllers.CreateCardHandler).
		Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(appURL, r))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
