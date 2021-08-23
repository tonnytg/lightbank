package migrations

import (
	"github.com/tonnytg/lightbank/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=lightbank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "test1", Email: "test1@test1.com"},
		{Username: "test2", Email: "test2@test1.com"},
	}
	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" +
			" account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})

	createAccounts()
}
