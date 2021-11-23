package migrations

import (
	"github.com/tonnytg/lightbank/database"
	"github.com/tonnytg/lightbank/helpers"
	"github.com/tonnytg/lightbank/interfaces"
)

func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "test1", Email: "test1@test1.com"},
		{Username: "test2", Email: "test2@test1.com"},
	}
	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" +
			" account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

func CreateUserAccounts() {
	database.DB.AutoMigrate(&interfaces.User{}, &interfaces.Account{})

	createAccounts()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}

// func MigrateTransaction() {
// 	Transaction := &interfaces.Transaction{}

// 	database.DB.AutoMigrate(&Transaction)
// }
