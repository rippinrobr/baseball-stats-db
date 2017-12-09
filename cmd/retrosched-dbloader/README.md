# retrosched-dbloader
You can load the database yourself if you have downloaded the [Retrosheet.org's Schedule Files](http://www.retrosheet.org/schedule/index.html) CSV files, loaded one of the schema files if you are going to use a relational database, and downloaded the latest `retrosched-dbloader` from the [latest release](https://github.com/rippinrobr/baseball-stats-db/releases). Here are the command line options available:

```
Usage of ./bin/retrosched-dbloader:
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
    	indicates what type of database is the load target. Supported databases are MongoDB, MySQL, Postgres, and SQLite
  -dbuser string
    	the username to use when loading the database. Required for all dbtypes except SQLite
  -inputdir string
    	the directory where the Retrosheet.org Schedule CSV files live. Required
  -verbose
    	writes more lines to the logs
```

#### Loading a MySQL DB

`./bin/retrosched-dbloader -dbtype mysql -dbname retrosheet -dbuser myusername -dbpass mypassword -inputdir ~/my-retrosheet-schedules`

This will load the data into your `retrosheet` database stored on the db server that lives on localhost since a `-dbhost` value wasn't provided.  Since the `-dbport` option was not provided the connection will attempt to use the default Postgres port 5432. 

#### Loading a Postgres DB

`./bin/retrosched-dbloader -dbtype postgres -dbname retrosheet -dbuser myusername -dbpass mypassword -inputdir ~/my-retrosheet-schedules`

This will load the data into your `retrosheet` database stored on the db server that lives on localhost since a `-dbhost` value wasn't provided.  Since the `-dbport` option was not provided the connection will attempt to use the default Postgres port 5432. 

#### Loading a SQLite DB

`./bin/retrosched-dbloader -dbtype sqlite -dbpath=./retrosheet.sqlite3 -inputdir ~/my-retrosheet-schedules`

Since this is a SQLite database there are only two required parameters, `-dbtype` and `-dbpath`.  The loader will create the SQLite database using the value of `-dbpath`.

## Licensing & Acknowledgements

```
The information used here was obtained free of
charge from and is copyrighted by Retrosheet.  Interested
parties may contact Retrosheet at 20 Sunset Rd., Newark, DE 19711.
```
