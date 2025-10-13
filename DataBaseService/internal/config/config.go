package config

import "os"

type Config struct {
	DatabaseURL    string
	GRPCPort       string
	MigrationsPath string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL:    getEnv("DATABASE_URL", "postgresql://postgres:password@localhost:5432/checklistdb"),
		GRPCPort:       getEnv("GRPC_PORT", ":50051"),
		MigrationsPath: getEnv("MIGRATIONS_PATH", "file://migrations"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
