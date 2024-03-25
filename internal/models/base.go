package models

import (
	database "github.com/christo-andrew/maybe-go/pkg/database"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Account{},
		&BankAccount{},
		&Transaction{},
		&CreditCardAccount{},
		&RealEstateAccount{},
		&Category{},
		&User{},
	)
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database.GetDB(database.DatabaseConfig{})
}
