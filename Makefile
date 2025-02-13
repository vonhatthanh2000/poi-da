APP_NAME = server
GOOSE_DRIVER='postgres'
GOOSE_DBSTRING='postgres://postgres:password@localhost:5432/layerg-da?sslmode=disable'
GOOSE_MIGRATION_DIR='./db/migrations'

build:
	go build -ldflags -w

api:
	go build -ldflags -w
	chmod +x layerg-da
	./layerg-da api --config .layerg-da.yaml

migrate-create:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql
migrate-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up
migrate-up-by-one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up-by-one
migrate-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) down
migrate-reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) reset	

seed:
	go run ./db/seeds/*.seed.go

sqlgen:
	sqlc generate