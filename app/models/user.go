package models

import (
	"time"
)

type UserRole string

const (
	UserRoleMerchant UserRole = "merchant"
	UserRoleCustomer UserRole = "customer"
)

type User struct {
	Email       *string    `json:"email" gorm:"email;primaryKey"`
	Password    *string    `json:"password" gorm:"password"`
	Role        *UserRole  `json:"role" gorm:"role"`
	FirstName   *string    `json:"firstName" gorm:"first_name"`
	LastName    *string    `json:"lastName" gorm:"last_name"`
	PhoneNumber *string    `json:"phoneNumber" gorm:"phone_number"`
	Address     *string    `json:"address" gorm:"address"`
	Cards       []Card     `json:"cards" gorm:"foreignKey:UserID"`
	CreatedDate *time.Time `json:"createdDate" gorm:"created_date"`
	UpdatedDate *time.Time `json:"updatedDate" gorm:"updated_date"`
}
