package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/bank/db/repository"
	"github.com/JoelD7/deuna-challenge/bank/models"
	"github.com/JoelD7/deuna-challenge/bank/usecases"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	cardNumberMin = 1000000000000000
	cardNumberMax = 9999999999999999
	ccvMin        = 100
	ccvMax        = 999
)

var (
	cardExpiryRegex = regexp.MustCompile(`^(0[1-9]|1[0-2])\/(0[0-9]|1[0-9]|2[0-9])$`)
)

func GetCardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardID, err := strconv.Atoi(vars["cardID"])
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

	getCard := usecases.NewCardGetter(repository.NewSQLiteClient())

	card, err := getCard(r.Context(), int64(cardID))
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	writeJSONData(w, card)
}

func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {
	cardValidator := usecases.NewCardValidator(repository.NewSQLiteClient())

	card, err := validateCard(r)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	err = cardValidator(r.Context(), card)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateCard(r *http.Request) (*models.Card, error) {
	var card models.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		return nil, err
	}

	if err = validateCardFields(&card); err != nil {
		return nil, err
	}

	return &card, nil
}

func validateCardFields(card *models.Card) error {
	if card.CardNumber == 0 {
		return models.ErrMissingCardNumber
	}

	if card.CardNumber < cardNumberMin || card.CardNumber > cardNumberMax {
		return models.ErrInvalidCardNumber
	}

	if card.Expiration == nil {
		return models.ErrMissingCardExpiration
	}

	if !cardExpiryRegex.MatchString(*card.Expiration) {
		return models.ErrInvalidCardExpirationFormat
	}

	if err := parseCardExpiration(*card.Expiration); err != nil {
		return err
	}

	if card.CCV == 0 {
		return models.ErrMissingCCV
	}

	if card.CCV < ccvMin || card.CCV > ccvMax {
		return models.ErrInvalidCCV
	}

	if card.Type == nil {
		return models.ErrMissingCardType
	}

	if *card.Type != models.CardTypeDebit && *card.Type != models.CardTypeCredit {
		return models.ErrInvalidCardType
	}

	return nil
}

func parseCardExpiration(expr string) error {
	date, err := time.Parse("01/06", expr)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErrInvalidCardExpirationDate, err)
	}

	date = date.AddDate(0, 0, 1-date.Day())

	if date.Before(time.Now()) {
		return models.ErrInvalidCardExpirationDate
	}

	return nil
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
