BIN := databank-dbloader
MAIN := cmd/databank-dbloader/main.go
RELEASE_DIR := release/
RCDB := baseball_databank_2017_21bf10f.sqlite3
TAG := 0.4.2
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

release:
	-mkdir $(RELEASE_DIR)

pkg-release-linux: release
	tar -zcvf $(RELEASE_DIR)$(LINUX_TGZ) $(RCDB) ./bin
	sha256sum $(RELEASE_DIR)$(LINUX_TGZ) >./$(RELEASE_DIR)$(LINUX_TGZ).checksum

pkg-release-macos: release
	tar -zcvf $(RELEASE_DIR)$(MACOS_TGZ) $(RCDB) ./bin
	shasum -a 256 $(RELEASE_DIR)$(MACOS_TGZ) >./$(RELEASE_DIR)$(MACOS_TGZ).checksum

release-linux: build pkg-release-linux

release-macos: build pkg-release-macos

