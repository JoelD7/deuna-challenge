package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/db/repository"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"net/http"
	"regexp"
	"time"
)

const (
	cardNumberMin = 1000000000000000
	cardNumberMax = 9999999999999999
	ccvMin        = 100
	ccvMax        = 999

	bankURL = "localhost:8081"
)

var (
	cardExpiryRegex = regexp.MustCompile(`^(0[1-9]|1[0-2])\/(0[0-9]|1[0-9]|2[0-9])$`)
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

	w.WriteHeader(http.StatusCreated)
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

	resp, err := http.Post(bankURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, models.ErrInvalidCard
	}

	if card.CardNumber == 0 {
		return nil, models.ErrMissingCardNumber
	}

	if card.CardNumber < cardNumberMin || card.CardNumber > cardNumberMax {
		return nil, models.ErrInvalidCardNumber
	}

	if card.Expiration == nil {
		return nil, models.ErrMissingCardExpiration
	}

	if !cardExpiryRegex.MatchString(*card.Expiration) {
		return nil, models.ErrInvalidCardExpirationFormat
	}

	if err = parseCardExpiration(*card.Expiration); err != nil {
		return nil, err
	}

	if card.CCV == 0 {
		return nil, models.ErrMissingCCV
	}

	if card.CCV < ccvMin || card.CCV > ccvMax {
		return nil, models.ErrInvalidCCV
	}

	if card.Type == nil {
		return nil, models.ErrMissingCardType
	}

	if *card.Type != *models.Debit && *card.Type != *models.Credit {
		return nil, models.ErrInvalidCardType
	}

	return &card, nil
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
