OS := $(shell uname)
DATABANK_BIN := databank-dbloader
DATABANK_MAIN := cmd/databank-dbloader/main.go
GL_BIN := retrogl-dbloader
GL_MAIN := cmd/retrogl-dbloader/main.go
SCHED_BIN := retrosched-dbloader
SCHED_MAIN := cmd/retrosched-dbloader/main.go
RELEASE_DIR := release/
HASH := NIGHTLY
INC_VERSION := 5
VERSION := 2017.$(INC_VERSION)
RCDB := baseball_databank_$(VERSION)_$(HASH).sqlite3
LINUX_TGZ := baseball-stats-db-$(VERSION)-linux-amd64.tgz
MACOS_TGZ := baseball-stats-db-$(VERSION)-macos-amd64.tgz
RELEASE_TGZ := baseball-stats-db-$(VERSION).tgz

all: lint vet test build_all

databank_build: $(DATABANK_MAIN) 
	go build -o bin/$(DATABANK_BIN) $(DATABANK_MAIN)

gl_build: $(GL_MAIN) 
	go build -o bin/$(GL_BIN) $(GL_MAIN)

sched_build: $(SCHED_MAIN) 
	go build -o bin/$(SCHED_BIN) $(SCHED_MAIN)

build_all: databank_build sched_build gl_build

vet: 
	go vet -all ./internal/... ./cmd/databank-dbloader/... ./cmd/retrosched-dbloader/... ./cmd/retrogl-dbloader/...

test: 
	go test ./...

lint:
	golint ./internal/... ./cmd/databank-dbloader/... ./cmd/retrosched-dbloader/... ./cmd/retrogl-dbloader/...

release_dir:
	-mkdir $(RELEASE_DIR)

sqlitedb:
	./bin/databank-dbloader -dbtype sqlite -dbpath=baseball_databank_2017.sqlite3 -inputdir ~/src/baseballdatabank/core 
	sqlite3 baseball_databank_2017.sqlite3 .schema >./schemas/sqlite_schema_$(VERSION).sql
	sqlite3 baseball_databank_2017.sqlite3 .backup >./backups/sqlite_backup$(VERSION).sql
	#cp  baseball_databank_2017.sqlite3 backups/$(RCDB)

mongodb: 
	./bin/databank-dbloader -dbtype mongodb -dbname baseball_databank -inputdir ~/src/baseballdatabank/core
	mongodump --archive=./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.archive --db baseball_databank
	tar zcvf ./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.tgz ./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.archive

mysqldb:  
	./bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core

postgresdb: 
	/bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core
	pg_dumpall >./backups/postgres_backup_$(VERSION)_$(HASH)_baseballdatabank.sql

pkg-release-linux: sqlitedb release_dir
	@echo "prepping release $(VERSION)"
	rm release/*
	tar -zcvf $(RELEASE_DIR)$(LINUX_TGZ) ./backups ./schemas
	sha256sum $(RELEASE_DIR)$(LINUX_TGZ) >./$(RELEASE_DIR)$(LINUX_TGZ).checksum

pkg-release-macos: sqlitedb release_dir
	rm release/*
	tar -zcvf $(RELEASE_DIR)$(MACOS_TGZ) ./bin ./backups
	shasum -a 256 $(RELEASE_DIR)$(MACOS_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum

release-linux: build pkg-release-linux

release-macos: build pkg-release-macos

release: release_dir #qlitedb
	@echo "Prepping release $(VERSION) $(OS)"
	rm $(RELEASE_DIR)*
	tar -zcvf $(RELEASE_DIR)$(RELEASE_TGZ) ./bin ./backups
ifeq ($(OS),Darwin) 
	shasum -a 256 $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
else
	sha256sum $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
endif
