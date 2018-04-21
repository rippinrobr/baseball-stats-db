DATABANK_BIN := databank-dbloader
DATABANK_MAIN := cmd/databank-dbloader/main.go
GL_BIN := retrogl-dbloader
GL_MAIN := cmd/retrogl-dbloader/main.go
SCHED_BIN := retrosched-dbloader
SCHED_MAIN := cmd/retrosched-dbloader/main.go
OS := $(shell uname)
SEASON := 2017
INC_VERSION := 5
VERSION := $(SEASON).$(INC_VERSION)
BDDB := baseball_databank_$(SEASON).sqlite3
RETRODB := retrosheet_$(SEASON).sqlite3
STATSDB := baseball_stats_$(SEASON).sqlite3 
RELEASE_DIR := release/
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

sqlitedb_bd:
	@echo "loading baseballdatabank db"
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(BDDB) -inputdir ~/src/baseballdatabank/core 
	sqlite3 $(BDDB) .schema >./schemas/sqlite_databank_schema_$(VERSION).sql
	sqlite3 $(BDDB) .dump >./backups/sqlite_databank_backup$(VERSION).sql
	
	@echo "loading baseball_stats db"
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(STATSDB) -inputdir ~/src/baseballdatabank/core 
	sqlite3 $(STATSDB) .schema >./schemas/sqlite_combined_schema_$(VERSION).sql
	sqlite3 $(STATSDB) .dump >./backups/sqlite_combined_backup$(VERSION).sql

sqlitedb_retro_stats_db:
	@echo "loading the baseball_stats db"
	./bin/retrosched-dbloader -dbtype sqlite -dbpath=./$(STATSDB) -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype sqlite -dbpath=./$(STATSDB) -inputdir ~/src/retrosheet/gamelog
	sqlite3 $(STATSDB) .schema >./schemas/sqlite_combined_schema_$(VERSION).sql 
	sqlite3 $(STATSDB) .dump >./backups/sqlite_combined_backup$(VERSION).sql 

sqlitedb_retro: sqlitedb_retro_stats_db
	@echo "loading the retrosheet db"
	./bin/retrosched-dbloader -dbtype sqlite -dbpath=./$(RETRODB) -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype sqlite -dbpath=./$(RETRODB) -inputdir ~/src/retrosheet/gamelog
	sqlite3 $(RETRODB) .schema >./schemas/sqlite_retrosheet_schema_$(VERSION).sql 
	sqlite3 $(RETRODB) .dump >./backups/sqlite_retrosheet_backup$(VERSION).sql 

sqlitedb: sqlitedb_bd sqlitedb_retro

mongodb: 
	./bin/databank-dbloader -dbtype mongodb -dbname baseball_databank -inputdir ~/src/baseballdatabank/core
	mongodump --archive=./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.archive --db baseball_databank
	tar zcvf ./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.tgz ./backups/mongodb_backup_$(VERSION)_$(HASH)_baseballdatabank.archive

mysqldb:  
	./bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core

postgresdb: 
	/bin/databank-dbloader --dbtype postgres --dbname baseballdatabank --dbuser postgres --dbpass itsmerob -inputdir ~/src/baseballdatabank/core
	pg_dumpall >./backups/postgres_backup_$(VERSION)_$(HASH)_baseballdatabank.sql

release: release_dir build_all sqlitedb
	@echo "Prepping release $(VERSION) $(OS)"
	rm $(RELEASE_DIR)*
	tar -zcvf $(RELEASE_DIR)$(RELEASE_TGZ) ./schemas ./backups
ifeq ($(OS),Darwin) 
	shasum -a 256 $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
else
	sha256sum $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
endif

release_all: build_all sqlitedb mysqldb postgresdb mongodb