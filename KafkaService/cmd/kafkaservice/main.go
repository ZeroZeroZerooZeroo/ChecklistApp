package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/kafkaservice/internal/config"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/kafkaservice/internal/kafka"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/kafkaservice/internal/logger"
)

func main() {
	cfg := config.LoadConfig()

	if err := logger.SetupLogger(cfg.LogFilePath); err != nil {
		log.Fatalf("Failed to setup logger: %v", err)
	}

	consumer, err := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	log.Printf("Kafka service started. Listening to topic: %s", cfg.KafkaTopic)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go consumer.ProcessMessages(ctx)

	<-sigchan
	log.Println("Shutting down Kafka service")

	cancel()
}
