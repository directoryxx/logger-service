package worker

import (
	"context"
	"fmt"
	"log"
	"logger/infrastructure"
	"logger/internal/controller"
	"logger/internal/repository"
	"logger/internal/usecase"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func RunWorkerServiceLog() {
	log.Println("[INFO] App Mode : Worker")

	log.Println("[INFO] Loading Database")
	dbMongo, err := infrastructure.ConnectMongo()

	if err != nil {
		log.Fatalf("Could not initialize Mongo connection using client %s", err)
	}

	// Ping the primary
	if err := dbMongo.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("[INFO] MongoDB Connected and Pinged")

	defer dbMongo.Disconnect(context.TODO())

	log.Println("[INFO] Loading Kafka Consumer")
	kafkaConn, err := infrastructure.ConnectKafka()

	if err != nil {
		log.Fatalf("Could not initialize connection to Kafka %s", err)
	}

	defer kafkaConn.Close()

	log.Println("[INFO] Loading Repository")
	logRepo := repository.NewLogRepository(dbMongo)

	log.Println("[INFO] Loading Usecase")
	logUsecase := usecase.NewLogUseCase(logRepo)

	log.Println("[INFO] Loading Controller")
	logController := controller.NewLogController(logUsecase)

	var wg sync.WaitGroup

	wg.Add(1)

	go consumerKafka("logger", kafkaConn, logController, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

func consumerKafka(topic string, kafkaConsumer *kafka.Consumer, logController controller.LogController, wg *sync.WaitGroup) {
	kafkaConsumer.SubscribeTopics([]string{topic}, nil)

	defer wg.Done()

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := kafkaConsumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			logController.InsertWorker(context.TODO(), string(msg.Value))
		}
	}

	kafkaConsumer.Close()
}
