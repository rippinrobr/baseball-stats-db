BIN := databank-dbloader
MAIN := cmd/databank-dbloader/main.go
RCDB := baseball_databank_2017_a663145.sqlite3
TAG := 0.4.0
LINUX_TGZ := baseball-stats-db-$(TAG)-linux-amd64.tgz
MACOS_TGZ := baseball-stats-db-$(TAG)-macos-amd64.tgz

all: lint vet test build linux

build: $(MAIN) 
	go build -o bin/$(BIN) $(MAIN)

vet: 
	go vet -all ./internal/... ./cmd/databank-dbloader/...

test: 
	go test ./...

lint:
	golint ./internal/... ./cmd/databank-dbloader/...

pkg-release-linux:
	tar -zcvf $(LINUX_TGZ) $(RCDB) ./bin
	sha256sum $(LINUX_TGZ) >./$(LINUX_TGZ).checksum

pkg-release-macos:
	tar -zcvf $(MACOS_TGZ) $(RCDB) ./bin
	shasum -a 256 $(MACOS_TGZ) >./$(MACOS_TGZ).checksum

release-linux: build pkg-release-linux

release-macos: build pkg-release-macos

