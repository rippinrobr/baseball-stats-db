# Baseball Stats DB 

The goal of the project is to share the data provided by the [Baseball Databank files](https://github.com/chadwickbureau/baseballdatabank) and [Retrosheet.org](http://retrosheet.org) in a way that makes it easier for other interested people to work with the data using SQL.  Currently this project does that by supplying database backup and schema files for the three officially supported databases PostgreSQL, SQLite and MySQL.  

_If there is a database you'd like to see supported, please open an [Issue](https://github.com/rippinrobr/baseball-stats-db/issues/new) and I will look into adding it._ 

## My Philosophy 

I will always create a database that stores the data as is, in columns that match with the column headers.  For the Baseball Databank data, I have created a table for each file that contains the exact columns that exist in the files.   For the `retrosheet` database, I store all of the schedules in a single table but the data is not changed. That way you can interact with the data in a way that you may already be using.  I will also create a more [normalized](https://en.wikipedia.org/wiki/Database_normalization) version of the data, one that adhears more to the proper database structure.  The `retrosheet_events` database is an example of a more normalized database.

## Project Structure

The project is laid out by using the data source as the root directory with two child directories, `backups` and `schemas`. 
Each of these directories store exactly what their name implies.  These two directories can hold backups/schemas for multiple databases that were created using the given data source.  As an example, I have created two databases based on the Retrosheet data, one named `retrosheet` and the other named `retrosheet_events`, the backups for both databases will live in the same backups directory.  In the root of each of the datasource directory there will be one or more [Docker](https://www.docker.com/) files that are used to create docker containers with the data already loaded in the database.  _When I've created the docker images and pushed thme to docker hub I will add those links here_

### Backups & Schema File Generation

Database backups will run whenever an update from a datasource has occurred or a schema change has happened.  A new schema file will be generated whenever there is a schema change.  So the backup files and schema files may look out of sync because of their versions but remember that database backups will happen more often than new schema files being generated.

### File Naming Convention

To make things easier on all of us the backup and schema file names will all have the same pattern: 

`database-type_database-name_(schema|backup).<version>.sql`

Here's an example of the Baseball Databank SQLite version backup file, the first of the calendar year:

`sqlite_baseballdatabank_backup.2019.1.sql` The compressed file will have the name `sqlite_baseballdatabank_backup.2019.1.tgz`


## The Future

- [ ] Create a Postgres docker container with the `baseballdatabank` data pre-loaded and push it to [Docker Hub]()
- [ ] Create a Postgres docker container with the `retrosheet_events` data pre-loaded and push it to [Docker Hub]()
- [ ] Create a MySQL docker container with the `baseballdatabank` data pre-loaded and push it to [Docker Hub]()
- [ ] Create a MySQL docker container with the `retrosheet_events` data pre-loaded and push it to [Docker Hub]()
- [ ] Create a normalized version of the `baseballdatabank` data

## The Tools I use

I have created a couple of tools to parse the CSV files and Retrosheet specific files.  The tool I used to load the Baseball Databank data and the Retrosheet schedules and game log files is [csv-to](https://github.com/rippinrobr/csv-to).  The `csv-to` project can be used to parse and load any csv file data into a MySQL/MaraiDB, Postgres or SQLite databases.

The Retrosheet Events data is loaded using a specific tool I've created called [retrosheet-events](https://github.com/rippinrobr/retrosheet-events).  It is specifically made for parsing and loading the events data into a normalized database schema I created.

Feel free to try both tools yourself, I'm always looking for feedback.

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
