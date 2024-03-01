package main

import (
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	appHost = os.Getenv("APP_HOST")
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)

	r.HandleFunc("/payments", controllers.CreatePaymentHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/payments/{paymentID}", controllers.GetPaymentHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/payments/process", controllers.ProcessPaymentHandler).
		Methods(http.MethodPost)

	r.HandleFunc("/card", controllers.CreateCardHandler).
		Methods(http.MethodPost)

	fmt.Println("App server running on", appHost)
	log.Fatal(http.ListenAndServe(appHost, r))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
