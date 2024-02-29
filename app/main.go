package main

import (
	"github.com/JoelD7/deuna-challenge/app/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/payments/{paymentID}", controllers.GetPaymentHandler).
		Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
