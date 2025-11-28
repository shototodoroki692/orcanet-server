build:
	@go build -o bin/orcanet-server

run: build
	@./bin/orcanet-server

test:
	@go test -v ./...