package usecases

import (
	"context"
	"github.com/JoelD7/deuna-challenge/app/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	passwordCost = bcrypt.DefaultCost
)

type UserManager interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, email string) (*models.User, error)
}

func NewUserCreator(userManager UserManager) func(ctx context.Context, user *models.User) error {
	return func(ctx context.Context, user *models.User) error {
		unhashedPassword := *user.Password

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(unhashedPassword), passwordCost)
		if err != nil {
			return err
		}

		password := string(hashedPassword)

		user.Password = &password
		now := time.Now()
		user.CreatedDate = &now

		err = userManager.CreateUser(ctx, *user)
		if err != nil {
			return err
		}

		return nil
	}
}

func NewUserAuthenticator(userManager UserManager) func(ctx context.Context, email, password string) (*models.User, error) {
	return func(ctx context.Context, email, password string) (*models.User, error) {
		user, err := userManager.GetUser(ctx, email)
		if err != nil {
			return nil, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
		if err != nil {
			return nil, models.ErrWrongCredentials
		}

		return user, nil
	}
}
