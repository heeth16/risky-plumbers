# Risky Plumbers
API for managing risks, including retrieving, creating, and viewing individual risks implemented in Go.

# Usage
## Run Server
```bash
go build -o server cmd/risky-plumbers/main.go
./server -ip=127.0.0.1 -port=8080
```

Alternatively, you can install it as a binary with:
```bash
go install github.com/heeth16/risky-plumbers/cmd/risky-plumbers@latest
risky-plumbers -ip=127.0.0.1 -port=8080 # Starts the server
```

## Run Unit Tests
```bash
go test ./...
```

# API Requests
## 1. Retrieve a List of Risks
```bash
curl -X GET "http://127.0.0.1:8080/v1/risks" -H "Accept: application/json"
```

## 2. Create a New Risk
```bash
curl -X POST "http://127.0.0.1:8080/v1/risks" \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-d '{
  "state": "open",
  "title": "Risk Title Example",
  "description": "Description of the risk."
}'
```

## 3. Retrieve an Individual Risk
```bash
curl -X GET "http://127.0.0.1:8080/v1/risks/{id}" -H "Accept: application/json"
```
