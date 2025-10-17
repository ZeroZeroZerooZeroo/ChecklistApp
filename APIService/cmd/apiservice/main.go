package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/config"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/grpc"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/handlers"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/kafka"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	grpcClient, err := grpc.NewGRPCClient(cfg.GetGRPCAddress())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	fmt.Printf("Connected to gRPC server at %s\n", cfg.GetGRPCAddress())

	kafkaProducer, err := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)

	}
	defer kafkaProducer.Close()

	fmt.Printf("Connected to Kafka at %v, topic: %s\n", cfg.KafkaBrokers, cfg.KafkaTopic)

	// Инициализация хендлера с gRPC клиентом
	handler := handlers.NewHandler(grpcClient, kafkaProducer)

	router := routes.SetupRouters(handler)

	serverAddr := cfg.GetServerAddress()

	fmt.Printf("Server started on %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
