package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tonnytg/lightbank/adapter/broker/kafka"
	"github.com/tonnytg/lightbank/adapter/factory"
	"github.com/tonnytg/lightbank/adapter/presenter/transaction"
	"github.com/tonnytg/lightbank/usecase/process_transaction"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Start cmd")

	// TODO: List
	// 1. Create a new file named "main.go"
	// 2. Copy the contents of this file into the new file
	// 3. Run the new file
	// 4. Fix the code in the new file
	// 5. Run the new file
	// 6. Fix the code in the new file
	// 7. Run the new file
	// 8. Fix the code in the new file

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group1",
		"auto.offset.reset": "earliest",
	}
	topics := []string{"transaction"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}


}
