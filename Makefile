.PHONY: build-api build-cli run-api build-cli static-check

build-api: 
	go build -o audiofile-api cmd/api/main.go

build-cli: 
	go build -o audiofile-cli cmd/cli/main.go

run-api:
	go run cmd/api/main.go

static-check:
	golangci-lint run