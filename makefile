include .env

create_migration:
	migrate create -ext=sql -dir=migrations -seq init

migrate_up:
	migrate -path=migrations -database "postgresql://${PG_DATABASE_USER}:${PG_DATABASE_PASSWORD}@${PG_DATABASE_HOST}:${PG_DATABASE_PORT}/${PG_DATABASE_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=migrations -database "postgresql://${PG_DATABASE_USER}:${PG_DATABASE_PASSWORD}@${PG_DATABASE_HOST}:${PG_DATABASE_PORT}/${PG_DATABASE_DB}?sslmode=disable" -verbose down

.PHONY: create_migration migrate_up migrate_down