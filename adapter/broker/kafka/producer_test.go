package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tonnytg/lightbank/adapter/presenter/transaction"
	"github.com/tonnytg/lightbank/domain/entity/transactions"
	"github.com/tonnytg/lightbank/usecase/process_transaction"

	"testing"
)

func TestProducerPublish(t *testing.T) {

	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       transactions.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	// TODO: create test for valid json format
	//outputJson, _ := json.Marshal(expectedOutput)


	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test",
		"auto.offset.reset": "earliest",
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	if err != nil {
		t.Errorf("Error publishing message: %s", err)
	}

}