package controllers

import (
	"encoding/json"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/JoelD7/deuna-challenge/app/usecases"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user, err := validateLoginRequest(r)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	authenticateUser := usecases.NewUserAuthenticator(sqliteClient)

	_, err = authenticateUser(r.Context(), *user.Email, *user.Password)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateLoginRequest(r *http.Request) (*models.User, error) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Email == nil {
		return nil, models.ErrMissingEmail
	}

	if user.Password == nil {
		return nil, models.ErrMissingPassword
	}

	return &user, nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user, err := validateSignupRequest(r)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	createUser := usecases.NewUserCreator(sqliteClient)

	err = createUser(r.Context(), user)
	if err != nil {
		models.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateSignupRequest(r *http.Request) (*models.User, error) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Email == nil {
		return nil, models.ErrMissingEmail
	}

	if user.Password == nil {
		return nil, models.ErrMissingPassword
	}

	if user.Role == nil {
		return nil, models.ErrMissingUserRole
	}

	if user.FirstName == nil {
		return nil, models.ErrMissingFirstName
	}

	if user.LastName == nil {
		return nil, models.ErrMissingLastName
	}

	if user.PhoneNumber == nil {
		return nil, models.ErrMissingPhoneNumber
	}

	if user.Address == nil {
		return nil, models.ErrMissingAddress
	}

	return &user, nil
}
