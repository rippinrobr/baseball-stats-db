# Baseball Stats DB 

The goal of the project is to share the data provided by the [Baseball Databank files](https://github.com/chadwickbureau/baseballdatabank) and [Retrosheet.org](http://retrosheet.org) in a format that allows users to easily load the data into a databases.  Currently there are three officially supported databases PostgreSQL, SQLite and MongoDB.  _MySQL isn't 'officially' support, [read this to find out why](cmd/databank-dbloader/README.md#about_mysql), but I do create backups and schema files for it._.  There are schema files and backups for a Baseball Databank only db, a Retrosheet only db plus a combined database named Baseball Stats.

## Schema Files
The schema files live in the  `schemas` directory. With the schema files for all the database will have the folowing naming convention.

`(mysql|postgres|sqlite)_(databank|retrosheet|combined)_schema_<season>.<version>.sql`

Season is the latest season's data that is contained in the Baseball Databank files.   The version number is the numbered
release for the season's data. A new version is released if the Baseball Databank or Retrosheet release an update to their
files.  Releases are typically released on Sundays.

## Backups
The backup files are included in each release.  Since the files can be big, I've removed the backups from the repository. The naming convention is similar to the schema files.  

`(mongodb|mysql|postgres|sqlite)_(databank|retrosheet|combined)_backup_<season>.<version>.tgz`

Season is the latest season's data that is contained in the Baseball Databank files.   The version number is the numbered
release for the season's data. A new version is released if the Baseball Databank or Retrosheet release an update to their
files.  Releases are typically released on Sundays.

## Utilities
I have removed the binaries from the release files, I believe that most people downloading the releases are there for the data.
I will be moving the code and binaries to my [sports-stats-utilities](https://github.com/rippinrobr/sports-stats-utilities), 
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
