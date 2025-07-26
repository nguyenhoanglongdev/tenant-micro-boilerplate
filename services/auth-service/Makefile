.PHONY: format lint test run

format:
	goimports -w .

lint:
	golangci-lint run

test:
	go test ./...

run:
	go run ./cmd/server/main.go
