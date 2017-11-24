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
_If you want to build the database yourself as of this writing you will need to create the data structures, database and build `dbloader` yourself.  The way the parsing is done relies on the path to the CSV file that was used to create the data structure which currently defaults to my environment.  I will be adding a command line option and changing the data structure generation script to fix this hard-coded requirement.  In the mean time you can use the backup files mentioned above._ 

Before running either of the python scripts you will need to install both `peewee` for creating the database schemas and  `inflection` (I used version 0.3.1 during development) for the data structure creation script. I installed the two packages using `pip`.

```
pip install peewee
pip install inflection
```

Once you have those two dependencies installed you are ready to run the scripts.


## Generating Data Structures
I start by creating the Go data structures that will be used to parse and store the data. There is a data structure created for each  [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) CSV files.  The script reads each of the CSV files, determines the data types of the columns and will write source code files for each of the CSV files. Currently only Go data structures are supported but in the future I will add others as I need them. If there is a particular language you would like to see you can create an issue or even better, submit a pull request.

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
  --name NAME           name of the data structure being created
  --output-dir OUTPUT_DIR
                        the directory where the generated file should be
                        written. If not provided file will be written to
                        stdout
  --verbose             more output during the parsing and creation of the
                        data structures
```

Before running this script you will need to have already downloaded repo.  In order to run the script you need to be in the `scripts/generate-code` directory. I ran the following to generate the models in this project:

`python create_data_structures.py --language go --csv --db --json --input-dir ~/src/baseballdatabank/core --package --output-dir ../../pkg/bd/models`

Required options: `--language` and one of the following `--input` or `--input-dir`.  `--input` takes a path to a single csv file while `--input-dir` takes a path to a directory of CSV files.

Since Go is the only supported language if `--language` is not provided it defaults to go.  This behavior will change as soon as another language is added. 

Optional parameters are `--name`, `--output-dir` and `--verbose`.  

The `--name` parameter is what will be used to name the generated data structure.  If `--name` is not provided the name of the file will be used to create the name following the language's naming conventions.  

The `--ouptut-dir` takes a path to the directory that the files should be written to.  If not provided the file(s) will be written to stdout.

#### Go specific command line options

There following arguments depend on language support `--csv`, `--db`, and `--json`. If any of the `--csv`, `--db`, and `--json` flags are used the generated structs will contain the `csv`, `db` and/or `json` tags.



### create_database_schema.py
Order of the data structure creation and schema script doesn't really matter but I typically create the schema after creating the structures.  Create the schema is as simple as running `python create_database_schema.py`.  The script lives in the `scripts/create-database` directory and has the options listed below.

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

Before you can use this script you will need to have already created the database in Postgres.  In future versions the database will be created for you.

#### Creating the Schema in a SQLite DB
`python create_database_schema.py --dbtype sqlite --dbpath ./baseball_database_db.py`

For SQLite any of the other command line arguments given besides the two used above will be ignored

#### Creating the Schema in a Postgres DB
`python create_database_schema.py --dbtype postgres --dbname baseballdatabank --dbuser myusername --dbpass mypassword`

For PostgreSQL if `--dbport` is not given the default port 5432 is used. 

After creating the schema you will have a db with a matching table for each of the CSV files


### dbloader 
After creating the database schema and the data structures you are ready to build the `dbloader`.  In the base directory of the project there is a `Makefile` that allows you to build the binary for your specific platform.  I built mine for my Mac using `make mac`. As you might expect if you want to build on a linux or windows machine you'd swap the mac for either of those two options.  The binary will bin the ./bin directory and has the following command line argmuments.

(https://github.com/chadwickbureau/baseballdatabank) CSV files.  You will need to clone the repository locally before running the utility.  In a future version you will have the ability to parse the files directly from github.  The following lists out the command line options.  Official releases will be on the github repository and will be available for Linux, MacOS and Windows.

```
Usage of ./bin/dbloader:
  -dbhost string
    	the name or ip address of the database server. (default "localhost")
  -dbname string
    	the hame of the database to load
  -dbpass string
    	the password to use when loading the database. Required for all dbtypes except SQLite
  -dbpath string
    	the path to your SQLite database
  -dbport int
    	the port to use when connecting to the database the database. Required for all dbtypes except SQLite
  -dbtype string
    	indicates what type of database is the load target. Supported databases are SQLite
  -dbuser string
    	the username to use when loading the database. Required for all dbtypes except SQLite
  -inputdir string
    	the directory where the Baseball Databank CSV files live. Required
  -verbose
    	writes more lines to the logs
```


#### Loading a Postgres DB

`./bin/dbloader -dbtype postgres -dbname baseballdatabank -dbuser myusername -dbpass mypassword -inputdir ~/my-baseball-databank-dir/core`

This will load the data into your `baseballdatabank` database stored on the db server that lives on localhost since a `-dbhost` value wasn't provided.  Since the -dbport option was provided the connection will attempt to use teh default Postgres port 5432.  The `-dbpath` option is only valid for SQLite databases.  

#### Loading a SQLite DB

`./bin/dbloader -dbtype sqlite -dbpath=./baseball_databank_2016_4a64a55.sqlite3 -inputdir ~/my-baseball-databank-dir/core`

Since this is a SQLite database there are only two required parameters, `-dbtype` and `-dbpath`.  The loader will create the SQLite database using the value of `-dbpath`.


