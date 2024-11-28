mockgen:
	go generate -v ./...

tidy:
	go mod tidy

run:
	go run cmd/main.go