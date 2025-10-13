package main

import (
	"log"
	"net"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/config"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/handler"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/repository"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/service"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/pkg/database"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/proto/checklist/pb"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	
	db, err := database.NewDatabase(cfg.Database.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func() {
		log.Println("Closing database connection")
		db.Close()
	}()

	
	if err := database.RunMigrations(cfg.Database.GetMigrationConnectionString()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	repo := repository.NewRepository(db.DB)
	serv := service.NewService(repo)
	grpchandler := handler.NewGRPCHandler(serv)

	grpcServer := grpc.NewServer()
	pb.RegisterCheckListServiceServer(grpcServer, grpchandler)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("DataBaseService (gRPC server) started on port %s", cfg.GRPCPort)
	log.Printf("Database: %s", cfg.Database.GetDBConnectionString())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
