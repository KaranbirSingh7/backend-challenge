all: build run

build:
	go build -o bin

run:
	go run main.go

dep:
	go mod tidy