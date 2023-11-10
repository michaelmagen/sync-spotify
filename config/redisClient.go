package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ctx = context.Background()

var RedisClient = newRedisClient()

func newRedisClient() *redis.Client {
	opt, _ := redis.ParseURL(getEnv("REDIS_CONNECTION_URL", ""))
	client := redis.NewClient(opt)

	// Example connection check
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return client
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	log.Println("Loading .env file")
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
