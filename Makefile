BIN := databank-dbloader
MAIN := cmd/databank-dbloader/main.go
RELEASE_DIR := release/
HASH := 9f91176
TAG := 2017.4
RCDB := baseball_databank_$(TAG)_$(HASH).sqlite3
LINUX_TGZ := baseball-stats-db-$(TAG)-linux-amd64.tgz
MACOS_TGZ := baseball-stats-db-$(TAG)-macos-amd64.tgz

all: lint vet test build

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

sqlitedb:
	./bin/databank-dbloader -dbtype sqlite -dbpath=baseball_databank_2017.sqlite3 -inputdir ~/src/baseballdatabank/core 
	cp  baseball_databank_2017.sqlite3 backups/$(RCDB)

mongodb: 
	./bin/databank-dbloader -dbtype mongodb -dbname baseball_databank -inputdir ~/src/baseballdatabank/core
	mongodump --archive=./backups/mongodb_backup_$(TAG)_$(HASH)_baseballdatabank.archive --db baseball_databank
	tar zcvf ./backups/mongodb_backup_$(TAG)_$(HASH)_baseballdatabank.tgz ./backups/mongodb_backup_$(TAG)_$(HASH)_baseballdatabank.archive

mysqldb:  
	./bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core

postgresdb: 
	/bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core
	pg_dumpall >./backups/postgres_backup_$(TAG)_$(HASH)_baseballdatabank.sql

pkg-release-linux: sqlitedb release
	rm release/*
	tar -zcvf $(RELEASE_DIR)$(LINUX_TGZ) ./bin ./backups
	sha256sum $(RELEASE_DIR)$(LINUX_TGZ) >./$(RELEASE_DIR)$(LINUX_TGZ).checksum

pkg-release-macos: sqlitedb release
	rm release/*
	tar -zcvf $(RELEASE_DIR)$(MACOS_TGZ) ./bin ./backups
	shasum -a 256 $(RELEASE_DIR)$(MACOS_TGZ) >./$(RELEASE_DIR)$(MACOS_TGZ).checksum

release-linux: build pkg-release-linux

release-macos: build pkg-release-macos

