BINARY_NAME=aprendiz-postgres
MAIN_PKG=./cmd/aprendiz-postgres

.PHONY: build run tidy test

tidy:
	go mod tidy

build: tidy
	go build -o bin/$(BINARY_NAME) $(MAIN_PKG)

run: tidy
	go run $(MAIN_PKG)

test:
	go test ./...
