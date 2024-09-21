# Build the application
build:
	@go build -o bin/golang-webserver_test ./cmd/golang-webserver

# Run the application
run: build
	@./bin/golang-webserver_test

# Run tests
test:
	@go test -v ./...

# Clean built binary
clean:
	@rm -f bin/golang-webserver_test