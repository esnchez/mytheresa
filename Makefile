MIGRATIONS_PATH=./internal/migrations 
DB_PATH=postgres://admin:admin@localhost/catalog?sslmode=disable

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_PATH) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_PATH) down

.PHONY: test
test:
	@go test -v ./...