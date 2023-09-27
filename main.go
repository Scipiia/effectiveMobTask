package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/scipiia/effectivemobiletask/api"
	db "github.com/scipiia/effectivemobiletask/db/sqlc"
	"github.com/scipiia/effectivemobiletask/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Api.NewServer error %w", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

// func runProcessorKafka() {
// 	taskProcessor, err := worker.NewKafkaProduce(context.Background(), "localhost:9092", "FIO")
// 	if err != nil {
// 		fmt.Println("cannot run task processor in main.go %w", err)
// 	}

// 	err = taskProcessor.Produce(context.Background())
// 	if err != nil {
// 		log.Fatal("cannot task processor in main.go", err)
// 	}
// }

// func runConsumerKafka() worker.Data {
// 	taskConsumer, err := worker.NewKafkaConsume(context.Background(), "localhost:9092", "FIO", "my-group")
// 	if err != nil {
// 		fmt.Println("cannot run task processor in main.go %w", err)
// 	}

// 	d, err := taskConsumer.Consume(context.Background())
// 	if err != nil {
// 		log.Fatal("cannot task consumer in main.go", err)
// 	}

// 	return d
// }
