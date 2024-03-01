package main

import (
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/controllers"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	appHost = os.Getenv("APP_HOST")
)

func main() {
	r := mux.NewRouter()
	r.Use(headerMiddleware)
	r.Use(authMiddleware)

	r.HandleFunc("/signup", controllers.SignupHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/login", controllers.LoginHandler).
		Methods(http.MethodPost)

	r.HandleFunc("/payments", controllers.CreatePaymentHandler).
		Methods(http.MethodPost)
	r.HandleFunc("/payments/{paymentID}", controllers.GetPaymentHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/payments/process", controllers.ProcessPaymentHandler).
		Methods(http.MethodPost)

	r.HandleFunc("/refund", controllers.RefundPaymentHandler).
		Methods(http.MethodPost)

	r.HandleFunc("/card", controllers.CreateCardHandler).
		Methods(http.MethodPost)

	fmt.Println("App server running on", appHost)
	log.Fatal(http.ListenAndServe(appHost, r))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signup" || r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		cookie := r.Header.Get("Cookie")
		if !strings.Contains(cookie, "accessToken") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		cookieParts := strings.Split(cookie, ";")
		accessToken := strings.Split(cookieParts[0], "=")[1]

		verifyToken := usecases.NewTokenValidator()

		_, err := verifyToken(r.Context(), accessToken)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
