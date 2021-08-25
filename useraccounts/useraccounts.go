package useraccounts

import (
	"github.com/tonnytg/lightbank/helpers"
	"github.com/tonnytg/lightbank/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
}