package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisClient(connectionString string, ttl int) (*RedisClient, error) {
	opts, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis connection string: %w", err)
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Printf("Successfully connected to Redis (TTL: %d seconds)", ttl)

	return &RedisClient{
		client: client,
		ttl:    time.Duration(ttl) * time.Second,
	}, nil
}


