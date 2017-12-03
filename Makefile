BIN := databank-dbloader
MAIN := cmd/databank-dbloader/main.go

mac: $(MAIN) vet test
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN) $(MAIN) 

linux: $(MAIN) vet test
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

windows: $(MAIN) vet test
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)


vet: $(MAIN)
	go vet -all ./internal/platform/db ./internal/databank ./internal/platform/parsers/csv ./cmd/databank-dbloader

test: $(MAIN)
	go test ./...
