package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost   string
	ServerPort   string
	GRPCHost     string
	GRPCPort     string
	KafkaBrokers []string
	KafkaTopic   string
}

func LoadConfig() *Config {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Printf("Warning: No .env file found or error loading: %v", err)
	}

	return &Config{
		ServerHost:   getEnv("SERVER_HOST", "localhost"),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		GRPCHost:     getEnv("GRPC_HOST", "localhost"),
		GRPCPort:     getEnv("GRPC_PORT", "50051"),
		KafkaBrokers: getEnvAsSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
		KafkaTopic:   getEnv("KAFKA_TOPIC", "user-actions"),
	}

}

func (c *Config) GetGRPCAddress() string {
	return c.GRPCHost + c.GRPCPort
}

func (c *Config) GetServerAddress() string {
	return c.ServerHost + ":" + c.ServerPort
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValues []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValues
}
