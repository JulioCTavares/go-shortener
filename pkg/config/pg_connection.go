package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() *pgx.Conn {
	var envs = LoadEnv()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envs.PgHost, envs.PgPort, envs.PgUser, envs.PgPassword, envs.PgDB)

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to connect to database: \n", err)
	}

	DB = conn

	fmt.Println("Connected to database PG!")

	return conn
}
