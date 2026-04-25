# HTTP Server From Scratch

Building a custom HTTP server in Go without using any external HTTP libraries.

## Learning Objectives
- Handle raw TCP connections in Go
- Parse raw HTTP requests from network streams
- Build HTTP responses from scratch
- Implement routing and request handlers
- Handle concurrent connections with goroutines

## Project Structure
- `main.go`: Entry point, starts the TCP listener
- `server.go`: TCP server implementation, connection handling
- `request.go`: HTTP request parsing
- `response.go`: HTTP response construction

## Setup

```bash
# Create module
go mod init github.com/yourusername/http-server

# Run the server
go run main.go

# Test in browser
open http://localhost:8080
```
