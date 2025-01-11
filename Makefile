MIGRATION_DIR=$(PWD)/db/migrations
SQLITE_DB=sqlite3://$(PWD)/db/videoverse/videoverse.db


build:
	go mod download; CGO_ENABLED=1 go build -o bin/videoverse ./cmd/*.go

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(SQLITE_DB) -verbose up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(SQLITE_DB) -verbose down

sqlc:
	rm -rf db/videoverse/*.go; sqlc generate

.PHONY: build
