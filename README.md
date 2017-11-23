# Baseball Databank DB 

The goal of the project is to provide the data provided by the [Baseball Databank files](https://github.com/chadwickbureau/baseballdatabank) in databases to make using the data easier.  Currently there are two officially supported databases, PostgreSQL and SQLite.  MongoDB and MySQL support are planned for future releases. Users can load data into MySQL but given that I cannot store Inf in a float field I'm not officially supporting MySQL until I can come up with a way to represent infinity in MySQL that isn't a complete hack MySQL will be supported.  

## Schemas
The schema file for all the supported databases live in the `schemas` directory.  All schema files follow this naming convention:

 `(postgres|sqlite)_schema_<season>_<github-commit-hash>.sql`

Where `<season>` the most recent season in the databases `<githhub-commit-hash>` is the hash of the latest [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) repository update. The SQLite schema for the most recent commit, as of 2017-11-22, would be `sqlite_schema_2016_4a64a55.sql`.  

## Backups
If want the schema AND the data you'll find backup files in the `backups/` directory.  The backup file naming convention is:

`(postgres|sqlite)_backup_<date of backup>_<season>_<github-commit-hash>.tgz`

If I had a backup file done today would have the name `postgres_backup_20171122_2016_4a64a55.tgz`

## Building the databases yourself
If you want to build the database yourself you'll need to run the `create_database_schema.py` script.  In order to that you'll need to have the `peewee` package installed, assuming you already have Python installed.  You can install `peewee` by running `pip install peewee`. Once you've installed the package you are ready to create the database.


### create_database_schema.py
Create the schema is as simple as running `python create_database_schema.py`.  The script lives in the `scripts/create-database` directory and has the options listed below.

```
usage: create_database_schema.py [-h] --dbtype {mysql,postgres,sqlite}
                                 [--dbhost DBHOST] [--dbname DBNAME]
                                 [--dbpath DBPATH] [--dbpass DBPASS]
                                 [--dbport DBPORT] [--dbuser DBUSER]

Generates a DB schema based on the Baseball Databank csv files.

optional arguments:
  -h, --help            show this help message and exit
  --dbtype {mysql,postgres,sqlite}
                        the database type you'd like to generate the schema
                        for
  --dbhost DBHOST       host of the database server
  --dbname DBNAME       Name of the database where the tables are to be added.
                        REQUIRED if not sqlite
  --dbpath DBPATH       SQLITE ONLY - the path for the newly created database
  --dbpass DBPASS       The password for the user given in the --dbuser
                        option, ignored for SQLite
  --dbport DBPORT       The port the database servier is listeing on, ignored
                        for SQLite, defaults to appropriate value for server
                        type if not provided
  --dbuser DBUSER       username to use when creating the database, ignorred
                        for SQLite databases, REQUIRED for others.
  ```

_For all databases except SQLite the database referenced by the --dbname option must already exist_

  Since SQLite is currently the only supported database running the `create_database_schema.py` script really only requires one command line argument since `--dbtype` defaults to `SQLite`.  Here's how to create the database:

  `python create_database_schema.py --dbpath ./baseball_database_db.py`

  If you were to run the script this way when the script finished you would have a new SQLite database in the directory that you ran the script in.  If you pass a path to the script of a pre-existing database you will receive an error.  In future version of the script this will be handled more gracefully.

### dbloader 
Once you have a database setup the `dbloader` executable can be used to [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) CSV files and load the data into the specified database. 

```
Usage of ./bin/dbloader:
  -dbconn string
    	the connection string used to log into a database
  -targetdb string
    	indicates what type of database is the load target. Supported databases are SQLite (default "sqlite")
  -verbose
    	writes more lines to the logs
```

Since SQLite is the only supported database at the moment the command line options are limited.  The only one that is required is the `-dbconn` option and it requires a path to the database to load.   In the near future the -dbconn string value will be determined by the value of the `-targetdb` option.  When there are new releases there will be executables for Linux, MacOS and Windows.

  ## Generating Data Structures
  This script creates the data structures code that describes the data in each of the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) CSV files.  The script reads each of the CSV files, determines the data types of the columns and will write source code files for each of the CSV files. In order to run this script you will need to have the package `inflection` installed. The script has been tested using `inflection` 0.3.1.  As of now Go is the only supported language, C#, Python, JavaScript and other languages are planned.  If there is a particular language you would like to see you can create an issue or even better, submit a pull request.

  ### create_data_structures.py
  ```
usage: create_data_structures.py [-h] [--csv] [--db] [--input INPUT]
                                 [--input-dir INPUT_DIR] [--json] [--package]
                                 --language {go} [--name NAME]
                                 [--output-dir OUTPUT_DIR] [--verbose]

Generate language specific data structures that model each of the Baseball
Databank CSV files.

optional arguments:
  -h, --help            show this help message and exit
  --csv                 If the language supports add CSV tags or CSV
                        representation to the structure
  --db                  adds db tags to Go structs.
  --input INPUT         the path to the input CSV file
  --input-dir INPUT_DIR
                        the path to the directory that contains the CSV files
                        to parse
  --json                If the language supports add JSON tags or JSON
                        representation to the structure
  --package             sets the package to the correct name for Go structs
  --language {go}       create language specific data structures
  --name NAME           name of the datas tructure being created
  --output-dir OUTPUT_DIR
                        the directory where the generated file should be
                        written. If not provided file will be written to
                        stdout
  --verbose             more output during the parsing and creation of the
                        data structures
```

The required options are `--language` and either the `--input` or `--input-dir` flags.  Since Go is the only supported language now if `--language` is not provided it defaults to go.  This behavior will change as soon as another language is added. The `--input` flag expects a path to a single CSV file.  THe `--input-dir` expects a path to a directory that contains *.CSV files.  

Optional parameters are `--name`, `--output-dir` and `--verbose`.  The `--name` parameter is what will be used to name the generated data structure.  If `--name` is not provided the name of the file will be used to create the name following the language's naming conventions.  The `--ouptut-dir` takes a path to the directory that the files should be written to.

#### Go specific command line options

There following arguments depend on language support `--csv`, `--db`, and `--json`. If any of the `--csv`, `--db`, and `--json` flags are used the generated structs will contain the `csv`, `db` and/or `json` tags.

