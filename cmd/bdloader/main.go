/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"
	"github.com/rippinrobr/baseball-databank-tools/pkg/db"
)

func main() {
	// cmd line args
	var dbconn, dbtype string
	var verbose bool
	// since SQLite is the only supported db at the moment then I will default to it for now
	flag.StringVar(&dbconn, "dbconn", "", "the connection string used to log into a database")
	flag.StringVar(&dbtype, "targetdb", db.DBSQLite, "indicates what type of database is the load target. Supported databases are SQLite")
	flag.BoolVar(&verbose, "verbose", false, "writes more lines to the logs")
	flag.Parse()

	if !db.IsSupportedDB(dbtype) {
		flag.Usage()
		fmt.Printf("\n'%s' is not a supported database\n", dbtype)
		os.Exit(1)
	}

	if dbconn == "" {
		if strings.ToLower(dbtype) == db.DBSQLite {
			log.Println("No dbconn value given for SQLite DB, create the baseball_databank.sqlite3 db in this dir.")
			dbconn = db.DefaultSQLiteDBName
		} else {
			log.Fatalf("A -dbconn value is required for the database type '%s'\n", dbtype)
		}
	}

	conn, connErr := db.CreateConnection(dbtype, dbconn)
	if connErr != nil {
		log.Fatal("Connection error: " + connErr.Error())
	}
	repo := db.CreateRepo(conn)
	defer repo.CloseConn()

	for _, o := range models.GetTableObjects() {
		log.Printf("Loading the table %s\n", o.GetTableName())
		if verbose {
			log.Printf("File Path: %s\n", o.GetFilePath())
			log.Printf("File Name: %s\n", o.GetFileName())
		}
		log.Printf("Loaded %d rows into %s\n\n", -1, o.GetTableName())
		// Get the Type of the object here and create an array of pointers
		// of the object's type to pass into the CSV parser
	}
}
