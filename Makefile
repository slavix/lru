
run:
	@go run ./cmd/app/.

test:
	@go test ./internal/lru -cover
