package models

type Merchant struct {
	ID          string `json:"id" gorm:"id"`
	FirstName   string `json:"firstName" gorm:"first_name"`
	LastName    string `json:"lastName" gorm:"last_name"`
	Email       string `json:"email" gorm:"email"`
	PhoneNumber string `json:"phoneNumber" gorm:"phone_number"`
	Address     string `json:"address" gorm:"address"`
}
