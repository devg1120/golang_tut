package pubsub

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	return client
}
