build:
	go mod download; go build -o bin/videoverse ./cmd/*.go

.PHONY: build
