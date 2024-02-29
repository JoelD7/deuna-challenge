package main

import (
	"encoding/json"
	"fmt"
	"github.com/JoelD7/deuna-challenge/bank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db, err := gorm.Open(sqlite.Open("deuna-bank-db.sqlt"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var transaction models.Transaction

	err = db.Model(&models.Transaction{}).Preload(clause.Associations).First(&transaction, "1").Error

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	data, err := json.Marshal(transaction)
	fmt.Println("Result: ", string(data))
}
