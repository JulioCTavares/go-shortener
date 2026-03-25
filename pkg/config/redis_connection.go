package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	var envs = LoadEnv()

	connStr := fmt.Sprintf("%s:%s", envs.RedisHost, envs.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: connStr,
	})

	defer client.Close()

	RedisClient = client

	fmt.Println("Connected to Redis!")

	return client
}
