FILES_CHANGED := ""
DATABANK_BIN := databank-dbloader
DATABANK_MAIN := cmd/databank-dbloader/main.go
DB_USER?=""
DB_PASS?=""
GL_BIN := retrogl-dbloader
GL_MAIN := cmd/retrogl-dbloader/main.go
SCHED_BIN := retrosched-dbloader
SCHED_MAIN := cmd/retrosched-dbloader/main.go
OS := $(shell uname) SEASON := 2018
INC_VERSION := 01
SEASON := 2018
VERSION := $(SEASON).$(INC_VERSION)
BDDB := baseballdatabank
RETRODB := retrosheet_$(SEASON)
STATSDB := baseball_stats_$(SEASON)
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
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(BDDB).sqlite3 -inputdir ~/src/baseballdatabank/core 
	sqlite3 $(BDDB).sqlite3 .schema >./schemas/sqlite_databank_schema_$(VERSION).sql
	sqlite3 $(BDDB).sqlite3 .dump >./backups/sqlite_databank_backup_$(VERSION).sql
	
	@echo "loading baseball_stats db"
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(STATSDB).sqlite3 -inputdir ~/src/baseballdatabank/core 
	sqlite3 $(STATSDB).sqlite3 .schema >./schemas/sqlite_combined_schema_$(VERSION).sql
	sqlite3 $(STATSDB).sqlite3 .dump >./backups/sqlite_combined_backup_$(VERSION).sql

sqlitedb_people:
	@echo "loading baseballdatabank db"
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(BDDB).sqlite3 -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	sqlite3 $(BDDB).sqlite3 .schema >./schemas/sqlite_databank_schema_$(VERSION).sql
	sqlite3 $(BDDB).sqlite3 .dump >./backups/sqlite_databank_backup_$(VERSION).sql
	
	@echo "loading baseball_stats db"
	./bin/databank-dbloader -dbtype sqlite -dbpath=$(STATSDB).sqlite3 -inputdir ~/src/baseballdatabank/core -inputfiles People.csv 
	sqlite3 $(STATSDB).sqlite3 .schema >./schemas/sqlite_combined_schema_$(VERSION).sql
	sqlite3 $(STATSDB).sqlite3 .dump >./backups/sqlite_combined_backup_$(VERSION).sql

sqlitedb_backups:
	@echo "backing up sqlitedbs"
	sqlite3 $(BDDB).sqlite3 .schema >./schemas/sqlite_databank_schema_$(VERSION).sql
	sqlite3 $(BDDB).sqlite3 .dump >./backups/sqlite_databank_backup_$(VERSION).sql
	sqlite3 $(STATSDB).sqlite3 .schema >./schemas/sqlite_combined_schema_$(VERSION).sql
	sqlite3 $(STATSDB).sqlite3 .dump >./backups/sqlite_combined_backup_$(VERSION).sql
	sqlite3 $(RETRODB).sqlite3 .schema >./schemas/sqlite_retrosheet_schema_$(VERSION).sql 
	sqlite3 $(RETRODB).sqlite3 .dump >./backups/sqlite_retrosheet_backup_$(VERSION).sql 


sqlitedb_retro_stats_db:
	@echo "loading the baseball_stats db"
	./bin/retrosched-dbloader -dbtype sqlite -dbpath=./$(STATSDB).sqlite3 -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype sqlite -dbpath=./$(STATSDB).sqlite3 -inputdir ~/src/retrosheet/gamelog
	sqlite3 $(STATSDB).sqlite3 .schema >./schemas/sqlite_combined_schema_$(VERSION).sql 
	sqlite3 $(STATSDB).sqlite3 .dump >./backups/sqlite_combined_backup$(VERSION).sql 

sqlitedb_retro: sqlitedb_retro_stats_db
	@echo "loading the retrosheet db"
	./bin/retrosched-dbloader -dbtype sqlite -dbpath=./$(RETRODB).sqlite3 -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype sqlite -dbpath=./$(RETRODB).sqlite3 -inputdir ~/src/retrosheet/gamelog
	sqlite3 $(RETRODB).sqlite3 .schema >./schemas/sqlite_retrosheet_schema_$(VERSION).sql 
	sqlite3 $(RETRODB).sqlite3 .dump >./backups/sqlite_retrosheet_backup$(VERSION).sql 

sqlitedb_tar: 
	tar zcvf ./backups/sqlite_databank_backup_$(VERSION).tgz ./backups/sqlite_databank_backup_$(VERSION).sql
	-tar zcvf ./backups/sqlite_retrosheet_backup_$(VERSION).tgz ./backups/sqlite_retrosheet_backup_$(VERSION).sql
	tar zcvf ./backups/sqlite_combined_backup_$(VERSION).tgz ./backups/sqlite_combined_backup_$(VERSION).sql
	
sqlitedb: sqlitedb_bd sqlitedb_retro sqlitedb_tar

mongodb_db: 
	./bin/databank-dbloader -dbtype mongodb -dbname $(BDDB) -inputdir ~/src/baseballdatabank/core
	./bin/databank-dbloader -dbtype mongodb -dbname $(STATSDB) -inputdir ~/src/baseballdatabank/core
	mongodump --archive=./backups/mongodb_databank_backup_$(VERSION).archive --db $(BDDB)
	mongodump --archive=./backups/mongodb_combined_backup_$(VERSION).archive --db $(STATSDB)

mongodb_people: 
	./bin/databank-dbloader -dbtype mongodb -dbname $(BDDB) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	./bin/databank-dbloader -dbtype mongodb -dbname $(STATSDB) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	mongodump --archive=./backups/mongodb_databank_backup_$(VERSION).archive --db $(BDDB)
	mongodump --archive=./backups/mongodb_combined_backup_$(VERSION).archive --db $(STATSDB)


mongodb_retro:
	./bin/retrosched-dbloader -dbtype mongodb -dbname $(RETRODB) -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype mongodb -dbname $(RETRODB) -inputdir ~/src/retrosheet/gamelog
	./bin/retrosched-dbloader -dbtype mongodb -dbname $(STATSDB) -inputdir ~/src/retrosheet/schedule
	./bin/retrogl-dbloader -dbtype mongodb -dbname $(STATSDB) -inputdir ~/src/retrosheet/gamelog
	mongodump --archive=./backups/mongodb_retrosheet_backup_$(VERSION).archive --db $(RETRODB)
	mongodump --archive=./backups/mongodb_combined_backup_$(VERSION).archive --db $(STATSDB)

mongodb_tar:
	tar zcvf ./backups/mongodb_databank_backup_$(VERSION).tgz ./backups/mongodb_databank_backup_$(VERSION).archive
	rm ./backups/mongodb_databank_backup_$(VERSION).archive
	-tar zcvf ./backups/mongodb_retrosheet_backup_$(VERSION).tgz ./backups/mongodb_retrosheet_backup_$(VERSION).archive
	-rm ./backups/mongodb_retrosheet_backup_$(VERSION).archive
	tar zcvf ./backups/mongodb_combined_backup_$(VERSION).tgz ./backups/mongodb_combined_backup_$(VERSION).archive
	rm ./backups/mongodb_combined_backup_$(VERSION).archive
	
mongodb: mongodb_db mongodb_retro mongodb_tar

mysqldb_db:  
	./bin/databank-dbloader --dbtype mysql --dbname $(BDDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core
	./bin/databank-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core

mysqldb_people:  
	./bin/databank-dbloader --dbtype mysql --dbname $(BDDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	./bin/databank-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv

mysqldb_retro:
	./bin/retrogl-dbloader --dbtype mysql --dbname $(RETRODB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrogl-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrosched-dbloader --dbtype mysql --dbname $(RETRODB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/schedule
	./bin/retrosched-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/schedule

mysqldb_retro_stats_db:
	./bin/retrogl-dbloader --dbtype mysql --dbname $(RETRODB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrogl-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrosched-dbloader --dbtype mysql --dbname $(RETRODB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/schedule
	./bin/retrosched-dbloader --dbtype mysql --dbname $(STATSDB) --dbuser $(DB_USER) --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/schedule

mysqldb_backup:
	mysqldump -u $(DB_USER) -p $(BDDB) >./backups/mysql_databank_backup_$(VERSION).sql
	mysqldump -u $(DB_USER) -p $(STATSDB) >./backups/mysql_combined_backup_$(VERSION).sql
	mysqldump -u $(DB_USER) -p $(RETRODB) >./backups/mysql_retrosheet_backup_$(VERSION).sql
	mysqldump --no-data -u $(DB_USER) -p $(RETRODB) >./schemas/mysql_retrosheet_schema_$(VERSION).sql
	mysqldump --no-data -u $(DB_USER) -p $(STATSDB) >./schemas/mysql_combined_schema_$(VERSION).sql
	mysqldump --no-data -u $(DB_USER) -p $(BDDB) >./schemas/mysql_databank_schema_$(VERSION).sql

mysqldb_tar: 
	tar zcvf ./backups/mysql_databank_backup_$(VERSION).tgz ./backups/mysql_databank_backup_$(VERSION).sql
	rm backups/mysql_databank_backup_$(VERSION).sql
	-tar zcvf ./backups/mysql_retrosheet_backup_$(VERSION).tgz ./backups/mysql_retrosheet_backup_$(VERSION).sql
	-rm backups/mysql_retrosheet_backup_$(VERSION).sql
	tar zcvf ./backups/mysql_combined_backup_$(VERSION).tgz ./backups/mysql_combined_backup_$(VERSION).sql
	rm backups/mysql_combined_backup_$(VERSION).sql

mysqldb: mysqldb_db mysqldb_retro #mysqldb_tar

postgresdb_db: 
	./bin/databank-dbloader --dbtype postgres --dbname $(BDDB) --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core
	./bin/databank-dbloader --dbtype postgres --dbname $(STATSDB) --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core
	pg_dump $(BDDB) >./backups/postgres_databank_backup_$(VERSION).sql
	pg_dump $(STATSDB) >./backups/postgres_combined_backup_$(VERSION).sql

postgresdb_people: 
	./bin/databank-dbloader --dbtype postgres --dbname $(BDDB) --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	./bin/databank-dbloader --dbtype postgres --dbname $(STATSDB) --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/baseballdatabank/core -inputfiles People.csv
	pg_dump $(BDDB) >./backups/postgres_databank_backup_$(VERSION).sql
	pg_dump $(STATSDB) >./backups/postgres_combined_backup_$(VERSION).sql

postgresdb_retro:
	./bin/retrogl-dbloader --dbtype postgres --dbname $(RETRODB) --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrosched-dbloader --dbtype postgres --dbname retrosheet_2017 --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrogl-dbloader --dbtype postgres --dbname baseball_stats_2017 --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	./bin/retrosched-dbloader --dbtype postgres --dbname baseball_stats_2017 --dbuser postgres --dbpass $(DB_PASS) -inputdir ~/src/retrosheet/gamelog
	pg_dump retrosheet_2017 >./backups/postgres_retrosheet_backup_$(VERSION).sql
	pg_dump baseball_stats_2017 >./backups/postgres_combined_backup_$(VERSION).sql

postgresdb_tar: 
	tar zcvf ./backups/postgres_databank_backup_$(VERSION).tgz ./backups/postgres_databank_backup_$(VERSION).sql
	rm backups/postgres_databank_backup_$(VERSION).sql
	-tar zcvf ./backups/postgres_retrosheet_backup_$(VERSION).tgz ./backups/postgres_retrosheet_backup_$(VERSION).sql
	-rm backups/postgres_retrosheet_backup_$(VERSION).sql
	tar zcvf ./backups/postgres_combined_backup_$(VERSION).tgz ./backups/postgres_combined_backup_$(VERSION).sql
	rm backups/postgres_combined_backup_$(VERSION).sql

postgresdb_schema:
	pg_dump --schema-only baseballdatabank >./schemas/postgres_databank_$(VERSION).sql 
	pg_dump --schema-only retrosheet_2017 >./schemas/postgres_retrosheet_$(VERSION).sql 
	pg_dump --schema-only baseball_stats_2017 >./schemas/postgres_combined_$(VERSION).sql 

postgresdb: postgresdb_db postgresdb_retro postgresdb_tar postgresdb_schema

databank: mysqldb_db mongodb_db postgresdb_db sqlitedb_db 
people: mongodb_people postgresdb_people sqlitedb_people
people_tar: mongodb_tar postgresdb_tar sqlitedb_tar
retrosheet: mysqldb_retro sqlitedb_retro mongodb_retro postgresdb_retro

release: release_dir 
	@echo "Prepping release $(VERSION) $(OS)"
	rm $(RELEASE_DIR)*
	tar -zcvf $(RELEASE_DIR)$(RELEASE_TGZ) README.md backups/*$(VERSION).tgz schemas/*$(VERSION).sql
ifeq ($(OS),Darwin) 
	shasum -a 256 $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
else
	sha256sum $(RELEASE_DIR)$(RELEASE_TGZ) >./$(RELEASE_DIR)$(RELEASE_TGZ).checksum
endif

release_all: build_all sqlitedb mysqldb postgresdb mongodb release
