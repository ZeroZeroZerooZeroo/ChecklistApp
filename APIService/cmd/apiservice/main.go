package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/grpc"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/handlers"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/routes"
)

func main() {
	// TODO: инициализация конфига

	//Временно пока без конфига
	grpcServerAddress := "localhost:50051"

	grpcClient, err := grpc.NewGRPCClient(grpcServerAddress)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	fmt.Printf("Connected to gRPC server at %s\n", grpcServerAddress)

	// Инициализация хендлера с gRPC клиентом
	handler := handlers.NewHandler(grpcClient)

	router := routes.SetupRouters(handler)

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
