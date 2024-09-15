build: 
	@go build -o bin/golang-webserver_test

run: build
	@./bin/golang-webserver_test

test: 
	@go test -v ./...