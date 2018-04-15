
# databank-dbloader 
You can load the database yourself if you have downloaded/cloned the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank) CSV files, loaded one of the schema files if you are going to use a relational database, and downloaded the latest `databank-dbloader` from the [latest release](https://github.com/rippinrobr/baseball-stats-db/releases). Here are the command line options available:

```
Usage of ./bin/databank-dbloader:
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
    	indicates what type of database is the load target. Supported databases are MongoDB, Postgres, and SQLite
  -dbuser string
    	the username to use when loading the database. Required for all dbtypes except SQLite
  -inputdir string
    	the directory where the Baseball Databank CSV files live. Required
  -verbose
    	writes more lines to the logs
```

#### Loading a Postgres DB

`./bin/databank-dbloader -dbtype postgres -dbname baseballdatabank -dbuser myusername -dbpass mypassword -inputdir ~/my-baseball-databank-dir/core`

This will load the data into your `baseballdatabank` database stored on the db server that lives on localhost since a `-dbhost` value wasn't provided.  Since the `-dbport` option was not provided the connection will attempt to use the default Postgres port 5432. 

#### Loading a SQLite DB

`./bin/databank-dbloader -dbtype sqlite -dbpath=./baseball_databank_2016_4a64a55.sqlite3 -inputdir ~/my-baseball-databank-dir/core`

Since this is a SQLite database there are only two required parameters, `-dbtype` and `-dbpath`.  The loader will create the SQLite database using the value of `-dbpath`.

<href name="about_mysql">
  
## About MySQL

The reason why MySQL isn't officially supported is due to the fact that there are 14 records that aren't in the `pitchingpost` table due to MySQL's decision to not support the `inf` value for infinity.  Having said that there is support in the `dbloader` code for MySQL.  If you choose to use MySQL just remember that you will not have all the data that is present in the CSV files.  Here are the missing records:

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
