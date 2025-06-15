__default:
    just --list

# Format flake and code
fmt:
    alejandra .
    golangci-lint run --fix ./...

# Run project locally
run:
    go run cmd/main/main.go
