# NetHTTP CRUD API

This is a simple REST API built with Go's `net/http` library. The API provides CRUD operations for managing Pokémon data.

## Getting Started

### Prerequisites

- Go 1.23.1 or later

### Installation

Clone the repository:
  ```sh
  git clone https://github.com/guyjof/nethttp-crud-api.git
  cd nethttp-crud-api
  ```

### Running the Application

To start the server, run:

```go
go run main.go
```

The server will be running at http://localhost:8080.

### API Endpoints
`GET /pokemon`: Get all Pokémon.

`GET /pokemon/{id}`: Get a Pokémon by ID.

`POST /pokemon`: Create a new Pokémon (requires authentication).

`GET /panic`: This route will intentionally cause a panic, which will be recovered by the RecoveryMiddleware.

### Middleware
`RecoveryMiddleware`: Recovers from panics and returns a 500 Internal Server Error.

`AuthMiddleware`: Handles authentication for the POST /pokemon route.

`LoggerMiddleware`: Logs incoming requests

### Error Handling
The RecoveryMiddleware is used to handle panics gracefully. It logs the error and stack trace, then returns a 500 Internal Server Error.
