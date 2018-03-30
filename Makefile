BIN := databank-dbloader
MAIN := cmd/databank-dbloader/main.go

all: lint vet test build

build: $(MAIN) 
	go build -o bin/$(BIN) $(MAIN)

vet: 
	go vet -all ./internal/... ./cmd/databank-dbloader/...

test: 
	go test ./...

lint:
	golint ./internal/... ./cmd/databank-dbloader/...