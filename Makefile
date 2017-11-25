BIN := dbloader
MAIN := cmd/dbloader/main.go

mac: $(MAIN) vet test
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN) $(MAIN) 

mac-release: mac
	goreleaser --rm-dist -c mac.goreleaser.yml

linux: $(MAIN) vet test
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

linux-release: linux
	goreleaser --rm-dist -c linux.goreleaser.yml

windows: $(MAIN) vet test
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

windows-release: windows
	goreleaser --rm-dist -c windows.goreleaser.yml

vet: $(MAIN)
	go vet -all ./pkg/db ./pkg/bd/models ./pkg/parsers/csv ./cmd/dbloader

test: $(MAIN)
	go test ./...
