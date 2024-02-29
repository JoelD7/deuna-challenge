package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/db/repository"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"io"
	"net/http"
	"os"
)

var (
	bankHost = os.Getenv("BANK_HOST")
)

func CreateCardHandler(w http.ResponseWriter, r *http.Request) {
	card, err := validateCreateCardRequest(r)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	createCard := usecases.NewCardCreator(repository.NewSQLiteClient())
	getCard := usecases.NewCardGetter(repository.NewSQLiteClient())

	id, err := createCard(r.Context(), card)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	newCard, err := getCard(r.Context(), id)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	writeJSONData(w, newCard)
}

func validateCreateCardRequest(r *http.Request) (*models.Card, error) {
	var card models.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		return nil, err
	}

	requestBody, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}

	url := "http://" + bankHost + "/card"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error sending request to bank: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusOK {
		return &card, nil
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("%w: %s", models.ErrInvalidCard, string(responseBody))
	}

	return nil, fmt.Errorf("error validating card: %s", string(responseBody))
}
