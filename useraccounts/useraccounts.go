package useraccounts

import (
	"fmt"
	"github.com/tonnytg/lightbank/domain/entity/transactions"

	"github.com/tonnytg/lightbank/database"
	"github.com/tonnytg/lightbank/helpers"
	"github.com/tonnytg/lightbank/interfaces"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAccounts := interfaces.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)

	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAccounts.ID = account.ID
	responseAccounts.Name = account.Name
	responseAccounts.Balance = int(account.Balance)
	return responseAccounts
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ? ", id).First(&account).Error != nil {
		return nil
	}

	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)

	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not owner of the account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}

		updatedAccount := updateAccount(from, int(fromAccount.Balance)-amount)
		updateAccount(to, int(toAccount.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = updatedAccount
		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
