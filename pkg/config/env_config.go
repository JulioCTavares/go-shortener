package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PgUser     string
	PgPassword string
	PgDB       string
	PgPort     string
	PgHost     string
	RedisPort  string
	RedisHost  string
}

var Envs *Config

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	Envs = &Config{
		PgUser:     os.Getenv("PG_DATABASE_USER"),
		PgPassword: os.Getenv("PG_DATABASE_PASSWORD"),
		PgDB:       os.Getenv("PG_DATABASE_DB"),
		PgPort:     os.Getenv("PG_DATABASE_PORT"),
		PgHost:     os.Getenv("PG_DATABASE_HOST"),
		RedisPort:  os.Getenv("REDIS_DATABASE_PORT"),
		RedisHost:  os.Getenv("REDIS_DATABASE_HOST"),
	}

	return Envs
}
