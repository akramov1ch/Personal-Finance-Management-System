CURRENT_DIR := $(shell pwd)
DB_URL := "postgres://postgres:azamat@localhost:5432/mydatabase?sslmode=disable"

migrate-up:
	migrate -path db/migrations -database $(DB_URL) up

migrate-down:
	migrate -path db/migrations -database  $(DB_URL) down

migrate-force:
	migrate -path db/migrations -database  $(DB_URL) force 1

run:
	@go run cmd/main.go