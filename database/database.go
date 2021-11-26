package database

import (
	"fmt"
	"github.com/tonnytg/lightbank/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfg struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func InitDatabase(cf *DBConfg) {
	dsn := fmt.Sprintf("host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		cf.Host, cf.Port, cf.SSLMode, cf.User, cf.Password, cf.Name)
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
