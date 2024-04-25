build:
	go build -o dsa-golang cmd/main.go

run:
	go run cmd/main.go

.PHONY: build run 
