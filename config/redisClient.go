package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
)

var ctx = context.Background()

// CustomRedisClient embeds *redis.Client and adds custom methods
type CustomRedisClient struct {
	*redis.Client
}

var RedisClient CustomRedisClient = newRedisClient()

func newRedisClient() CustomRedisClient {
	opt, _ := redis.ParseURL(getEnv("REDIS_CONNECTION_URL", ""))
	client := redis.NewClient(opt)

	// Example connection check
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return CustomRedisClient{client}
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

func (rc *CustomRedisClient) SetToken(ctx context.Context, token oauth2.Token) error {
	// Marshal the OAuth2 token to JSON
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// Set the JSON token in Redis with the access code as the key
	err = rc.Set(ctx, token.AccessToken, tokenJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
