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
