package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/config"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/grpc"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/handlers"
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

	// Инициализация хендлера с gRPC клиентом
	handler := handlers.NewHandler(grpcClient)

	router := routes.SetupRouters(handler)

	serverAddr := cfg.GetServerAddress()

	fmt.Printf("Server started on %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
