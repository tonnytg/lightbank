package transaction
// show kafka transaction
// this is presenter of clean architecture

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/lightbank/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (t *KafkaPresenter) Bind(input interface{}) error {
	fmt.Println("kafka aqui: 1")
	t.ID = input.(process_transaction.TransactionDtoOutput).ID
	t.Status = input.(process_transaction.TransactionDtoOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (t *KafkaPresenter) Show() ([]byte, error){
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}