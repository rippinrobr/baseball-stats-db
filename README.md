# Baseball Stats DB 

The goal of the project is to provide the data provided by the [Baseball Databank files](https://github.com/chadwickbureau/baseballdatabank)  in databases to make using the data easier.  Currently there are three officially supported databases PostgreSQL, SQLite and MongoDB.  _If you are wondering why MySQL isn't support [read this](cmd/databank-dbloader/README.md#about_mysql)_.  

Another source of data that will be a part of this database is from [Retrosheet.org](http://retrosheet.org).  The Retrosheet project has game logs and other data points that the Baseball Databank project does not.  Currently the data will be in two separate databases but future releases will include a combined database.

## Schemas
The schema file for all the supported databases live in the `schemas` directory. Each database will have a similar name convention.  The All schema files follow this naming convention:

### Baseball Databank
The [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) schema has the following naming convention:

 `(postgres|sqlite)_schema_<season>_<github-commit-hash>.sql`


Where `<season>` the most recent season in the databases `<githhub-commit-hash>` is the hash of the latest [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) repository update. The SQLite schema for the most recent commit, as of 2017-11-22, would be `sqlite_schema_2016_4a64a55.sql`.  
### Retrosheet 
The [Retrosheet.org](http://www.retrosheet.org) database schema has the following naming convention:

 `(mysql|postgres|sqlite)_schema_<season>_retrosheet.sql


## Backups
If want the schema AND the data you'll find backup files in the `backups/` directory.  The backup file naming convention is:

`(postgres|sqlite|mongodb)_backup_<date of backup>_<season>_<github-commit-hash>.tgz`

If I had a backup file done today would have the name `postgres_backup_20171122_2016_4a64a55.tgz`

## Utilities
In addition to the schema and backup files there are utilities that you can use to create load the databases yourself.

* `databank-dbloader` is used to load Baseball Databank into a database of your choosing. For more information see the [README.md](cmd/databank-dbloader/README.md)l
* `retsched-dbloader` is used to load Retrosheet's schedules file into a database of your choosing. For more information see its [README.md](cmd/retrosched-dbloader/README.md).

## Releases & Versioning
As of April 8th, 2018 I will use the following version structure.  Each version will be <Latest Season Stats in db>.<update number>
So the third release of of the 2017 stats would have the version number `2017.3`. The release notes will also contain links to the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) commits that were added to the db as part of this release.

## Licensing & Acknowledgments

```
Baseball Databank is a compilation of historical baseball data in a
convenient, tidy format, distributed under Open Data terms.

This work is licensed under a Creative Commons Attribution-ShareAlike
3.0 Unported License.  For details see:
http://creativecommons.org/licenses/by-sa/3.0/

Person identification and demographics data are provided by
Chadwick Baseball Bureau (http://www.chadwick-bureau.com),
from its Register of baseball personnel.

Player performance data for 1871 through 2014 is based on the
Lahman Baseball Database, version 2015-01-24, which is 
Copyright (C) 1996-2015 by Sean Lahman.

The tables Parks.csv and HomeGames.csv are based on the game logs
and park code table published by Retrosheet.
This information is available free of charge from and is copyrighted
by Retrosheet.  Interested parties may contact Retrosheet at 
http://www.retrosheet.org.
```

```
The information used here was obtained free of
charge from and is copyrighted by Retrosheet.  Interested
parties may contact Retrosheet at 20 Sunset Rd., Newark, DE 19711.
```
