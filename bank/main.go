package main

import (
	"github.com/JoelD7/deuna-challenge/bank/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	bankURL = os.Getenv("BANK_URL")
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)

	r.HandleFunc("/card", controllers.ValidateCardHandler).
		Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(bankURL, r))
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
