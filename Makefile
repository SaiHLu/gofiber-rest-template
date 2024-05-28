# Load .env file
include .env

.PHONY: build
build:
	@go build -o bin/main cmd/api/*.go

.PHONY: run
run: build
	@./bin/main

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: create-migration
create-migration:
	@migrate create -ext sql -dir database/migrations -format "unix" $(file-name)

.PHONY: run-migrations
run-migrations:
	@migrate -path database/migrations -database postgres://$(POSTGRESDB_USERNAME):$(POSTGRESDB_PASSWORD)@$(POSTGRESDB_HOST):$(POSTGRESDB_PORT)/postgres?sslmode=disable up

.PHONY: rollback-migrations
rollback-migrations:
	@migrate -path database/migrations -database postgres://$(POSTGRESDB_USERNAME):$(POSTGRESDB_PASSWORD)@$(POSTGRESDB_HOST):$(POSTGRESDB_PORT)/postgres?sslmode=disable down
	