__default:
    just --list

# Format flake and code
fmt:
    alejandra .
    golangci-lint run --fix ./...

run:
    go run cmd/main/main.go

worker:
    go run cmd/worker/main.go

temporal-up:
    temporal server start-dev
