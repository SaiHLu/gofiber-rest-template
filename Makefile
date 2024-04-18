.PHONY: build
build:
	@go build -o bin/main cmd/*.go

.PHONY: run
run:
	@go build -o bin/main cmd/*.go
	@./bin/main

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: create-migration
create-migration:
	@migrate create -ext sql -dir database/migrations -format "unix" $(file-name)

.PHONY: run-migrations
run-migrations:
	@migrate -path database/migrations -database postgres://postgres:pass123@localhost:5432/golang?sslmode=disable up

.PHONY: rollback-migrations
rollback-migrations:
	@migrate -path database/migrations -database postgres://postgres:pass123@localhost:5432/golang?sslmode=disable down
	
	