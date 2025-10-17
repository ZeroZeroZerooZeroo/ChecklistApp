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

func (r *RedisClient) SetTask(ctx context.Context, key string, task interface{}) error {
	err := r.client.Set(ctx, key, task, r.ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set task in Redis: %w", err)

	}
	return nil
}

func (r *RedisClient) GetTask(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key not found in Redis")

	} else if err != nil {
		return "", fmt.Errorf("failed to get task from Redis: %w", err)
	}
	return val, nil
}

func (r *RedisClient) DeleteTask(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete task from Redis: %w", err)
	}
	return nil
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
