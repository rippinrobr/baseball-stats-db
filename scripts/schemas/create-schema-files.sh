#!/bin/bash

DBTYPE=$1
DBCONN=$2
SQLFILE=$3

sqlite3 $DBCONN .schema >$SQLFILE