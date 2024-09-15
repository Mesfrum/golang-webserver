# Golang Web Server

This is a simple web server written in Go that handles JSON requests.

## Project Structure

```
golang-webserver/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── server.go
│   │   └── handlers.go
│   └── config/
│       └── config.go
├── pkg/
│   └── jsonutil/
│       └── jsonutil.go
├── tests/
│   └── api_test.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.16 or later
- Make (optional, for using the Makefile)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/golang-webserver.git
   ```
2. Navigate to the project directory:
   ```
   cd golang-webserver
   ```
3. Install dependencies:
   ```
   go mod download
   ```

## Usage

To build and run the server:

```
make run
```

Or without Make:

```
go build -o bin/server cmd/server/main.go
./bin/server
```

The server will start on `http://localhost:3000` by default.

To send a request to the server:

```
curl -X POST -H "Content-Type: application/json" -d '{"key": "value"}' http://localhost:3000
```

## Configuration

The server can be configured using environment variables:

- `PORT`: The port on which the server will listen (default: 3000)
- `LOG_LEVEL`: The log level (default: "info")

## Testing

To run tests:

```
make test
```

Or without Make:

```
go test ./...
```

## License

This project is licensed under the MIT License.