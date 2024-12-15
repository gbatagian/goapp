# GoWS

A sample Go WebSocket application 

# Run

## Server 

### Local

* Direct run: `make run` (or `go run cmd/server/main.go`)
* Or build and run the binary:
  - Build binaries: `make`
  - Start the server: `./bin/server`

### Docker

* Run the container: `make run-container`

## CLI Client

1. Direct run: `go run cmd/client/main.go -n 10`
2. Or build and run the binary:
   - Build binaries: `make`    
   - Run the client: `./bin/client -n 10`