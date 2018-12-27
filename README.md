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

## About MySQL

The reason why MySQL isn't officially supported is due to the fact that there are 14 records that aren't in the `pitchingpost` table. The aren't there due to MySQL's decision to not support the `inf` value for infinity.  Having said that there is support in the `dbloader` code for MySQL.  If you choose to use MySQL just remember that you will not have all the data that is present in the CSV files.

```
playerID,yearID,round,teamID,lgID,W,L,G,GS,CG,SHO,SV,IPouts,H,ER,HR,BB,SO,BAOpp,ERA,IBB,WP,HBP,BK,BFP,GF,R,SH,SF,GIDP
poledi01,1975,WS,BOS,AL,0,0,1,0,0,0,0,0,0,1,0,2,0,,inf,0,0,0,0,2,0,1,0,0,0
welchbo01,1981,WS,LAN,NL,0,0,1,1,0,0,0,0,3,2,0,1,0,1.000,inf,0,0,0,0,4,0,2,0,0,0
westda01,1991,WS,MIN,AL,0,0,2,0,0,0,0,0,2,4,1,4,0,1.000,inf,0,0,0,0,6,0,4,0,0,0
holtch01,1999,NLDS1,HOU,NL,0,0,1,0,0,0,0,0,3,3,0,0,0,1.000,inf,0,0,0,0,3,0,3,0,0,0
fultzaa01,2002,NLDS1,SFN,NL,0,0,2,0,0,0,0,0,2,1,0,0,0,1.000,inf,0,0,0,0,2,0,1,0,0,0
myersmi01,2006,ALDS1,NYA,AL,0,0,1,0,0,0,0,0,1,1,1,0,0,1.000,inf,0,0,0,0,1,0,1,0,0,0
saitota01,2008,NLDS1,LAN,NL,0,0,1,0,0,0,0,0,3,2,0,0,0,1.000,inf,0,0,0,0,3,0,2,0,0,0
ramirra02,2009,ALDS2,BOS,AL,0,0,1,0,0,0,0,0,1,2,0,1,0,1.000,inf,0,0,1,0,3,0,2,0,0,0
schleda01,2011,ALCS,DET,AL,0,0,1,0,0,0,0,0,1,1,0,0,0,1.000,inf,0,0,0,0,1,0,1,0,0,0
ueharko01,2011,ALDS2,TEX,AL,0,0,1,0,0,0,0,0,2,3,1,1,0,1.000,inf,0,0,0,0,3,0,3,0,0,0
marshse01,2013,NLWC,CIN,NL,0,0,1,0,0,0,0,0,1,1,0,2,0,1.000,inf,1,0,0,0,3,0,1,0,0,0
choatra01,2014,NLDS2,SLN,NL,0,0,1,0,0,0,0,0,1,1,1,0,0,1.000,inf,0,0,0,0,1,0,1,0,0,0
goedder01,2015,NLDS2,NYN,NL,0,0,1,0,0,0,0,0,4,3,1,0,0,1.000,inf,0,0,0,0,4,0,3,0,0,0
jimenub01,2016,ALWC,BAL,AL,0,1,1,0,0,0,0,0,3,3,1,0,0,1.000,inf,0,0,0,0,3,1,3,0,0,0
```

I do not want to put some placeholder value in place of the `inf` simply because MySQL doesn't support it when PostgreSQL and SQLite do.  If you can come up with a clever way of handling this for MySQL that isn't a complete hack feel free to pass it along or make a pull request.

## Makefile

If you wish to use this repo to parse and load the databases yourself, you should use the makefile.  In order to connect to a database you need to pass in the username and password for the database you are going to connect to.  For instance, If I wanted to load a postgres database I'd run the following:

`make DB_USER=<db username> DB_PASS=<db password> postgresdb_db`

## Future

In the future the utility used in this repository will be replaced by my [csv-to]() project.  I now use it for PostgreSQL and SQLite but it does not support MySQL at this point.

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
