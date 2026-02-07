
.PHONY: run test lint clean


run: 
	go run cmd/go-pacman-version/main.go

build: 
	go build -o bin/go-pacman-version cmd/go-pacman-version/main.go

test: 
	go test -v -race ./...

lint: 
	golangci-lint run

clean: 
	go clean
	rm -f bin/go-pacman-version

