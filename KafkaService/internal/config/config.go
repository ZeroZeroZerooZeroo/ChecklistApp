package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBrokers []string
	KafkaTopic   string
	KafkaGroupID string
	LogFilePath  string
}

func LoadConfig() *Config {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Printf("Warning: No .env file found or error loading: %v", err)
	}

	return &Config{
		KafkaBrokers: getEnvAsSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
		KafkaTopic:   getEnv("KAFKA_TOPIC", "user-actions"),
		KafkaGroupID: getEnv("KAFKA_GROUP_ID", "kafka-service"),
		LogFilePath:  getEnv("LOG_FILE_PATH", "./logs/user-actions.log"),
	}

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
