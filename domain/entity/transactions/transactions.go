package transactions

import (
	"errors"
	"github.com/tonnytg/lightbank/database"
	creditcard "github.com/tonnytg/lightbank/domain/entity/credit_card"
	"github.com/tonnytg/lightbank/helpers"
	"github.com/tonnytg/lightbank/interfaces"
)

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   creditcard.CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}
	return nil
}

func (t *Transaction) SetCreditCard(cc creditcard.CreditCard) {
	t.CreditCard = cc
}

func CreateTransaction(From, To uint, Amount int) {
	transaction := &interfaces.SimpleTransaction{From: From, To: To, Amount: Amount}
	database.DB.Create(transaction)
}

func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction {
	transactions := []interfaces.ResponseTransaction{}
	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(interfaces.SimpleTransaction{From: id}).Or(interfaces.SimpleTransaction{To: id}).Scan(&transactions)
	return transactions
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ? ", id).Scan(&accounts)

		transactions := []interfaces.ResponseTransaction{}
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
