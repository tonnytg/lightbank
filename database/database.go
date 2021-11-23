package database

import (
	"github.com/tonnytg/lightbank/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=lightbank port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.HandleErr(err)
	DB = database
}

//func ConnectDB() *gorm.DB {
//	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=lightbank port=5432 sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	helpers.HandleErr(err)
//	return db
//}
