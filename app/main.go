package main

import (
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("deuna-db.sqlt"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var payment models.Payment

	err = db.Model(&models.Payment{}).Preload("Card").First(&payment, "1").Error

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	data, err := json.Marshal(payment)
	fmt.Println("Result: ", string(data))
}
