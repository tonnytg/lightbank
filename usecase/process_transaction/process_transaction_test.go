package process_transaction

import (
	"github.com/golang/mock/gomock"
	mock_broker "github.com/tonnytg/lightbank/adapter/broker/mock"

	"github.com/tonnytg/lightbank/domain/entity/transactions"
	mock_repository "github.com/tonnytg/lightbank/domain/repository/mock"
	"testing"
	"time"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "1111111111111111",
		CreditCardName:            "Test Beta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)

	// TODO: Está com erro
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	//Publish(msg interface{}, key []byte, topic string) error
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}

func TestProcessTransaction_ExecuteRejectTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4001111111111111",
		CreditCardName:            "Test Besta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1200,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4311111111111111",
		CreditCardName:            "Test Besta",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    900,
	}
	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.APPROVED,
		ErrorMessage: "",
	}

	// manage mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// outputs must be equal
	if expectedOutput != output {
		t.Errorf("expected %v, got %v", expectedOutput, output)
	}
}
