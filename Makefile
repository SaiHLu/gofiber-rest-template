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

.PHONY: make-migration
make-migration:
	@migrate create -ext sql -dir database/migration -format "unix" $(name)

.PHONY: migrate
migrate:
	@migrate -path database/migration -database postgres://$(POSTGRESDB_USERNAME):$(POSTGRESDB_PASSWORD)@$(POSTGRESDB_HOST):$(POSTGRESDB_PORT)/postgres?sslmode=disable up $(steps)

.PHONY: migrate-rollback
migrate-rollback:
	@migrate -path database/migration -database postgres://$(POSTGRESDB_USERNAME):$(POSTGRESDB_PASSWORD)@$(POSTGRESDB_HOST):$(POSTGRESDB_PORT)/postgres?sslmode=disable down $(steps)
	