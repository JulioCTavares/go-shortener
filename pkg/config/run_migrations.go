package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	var env = LoadEnv()

	connStr := fmt.Sprint("postgresql://", env.PgUser, ":", env.PgPassword, "@", env.PgHost, ":", env.PgPort, "/", env.PgDB, "?sslmode=disable")

	m, err := migrate.New("file://migrations", connStr)

	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to apply migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
	return nil
}
