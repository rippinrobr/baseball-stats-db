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
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"strings"

	"github.com/rippinrobr/baseball-databank-db/pkg/bd/models"
	"github.com/rippinrobr/baseball-databank-db/pkg/db"
)

const (
	dbpathParam  string = "dbpath"
	dbhostParam         = "dbhost"
	dbnameParam         = "dbname"
	dbpassParam         = "dbpass"
	dbportParam         = "dbport"
	dbtypeParam         = "dbtype"
	dbuserParam         = "dbuser"
	verboseParam        = "verbose"
)

func main() {
	// cmd line args
	var dbhost, dbname, dbpass, dbpath, dbtype, dbuser string
	var dbport int
	var verbose bool
	// since SQLite is the only supported db at the moment then I will default to it for now
	flag.StringVar(&dbhost, dbhostParam, "localhost", "the name or ip address of the database server.")
	flag.StringVar(&dbname, dbnameParam, "", "the hame of the database to load")
	flag.StringVar(&dbpass, dbpassParam, "", "the password to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbpath, dbpathParam, "", "the path to your SQLite database")
	flag.IntVar(&dbport, dbportParam, 0, "the port to use when connecting to the database the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbtype, dbtypeParam, "", "indicates what type of database is the load target. Supported databases are SQLite")
	flag.StringVar(&dbuser, dbuserParam, "", "the username to use when loading the database. Required for all dbtypes except SQLite")
	flag.BoolVar(&verbose, verboseParam, false, "writes more lines to the logs")
	flag.Parse()

	dbtype = strings.ToLower(dbtype)
	if !db.IsSupportedDB(dbtype) {
		flag.Usage()
		fmt.Printf("\n'%s' is not a supported database\n", dbtype)
		os.Exit(1)
	}

	if dbtype == db.DBSQLite && dbpath == "" {
		log.Fatalf("A -dbpath value is required for the database type '%s'\n", dbtype)
	}

	processFiles(db.Options{
		Host:    dbhost,
		Name:    dbname,
		Pass:    dbpass,
		Path:    dbpath,
		Port:    dbport,
		Type:    dbtype,
		User:    dbuser,
		Verbose: verbose,
	})
}

func processFiles(opts db.Options) {
	if opts.Verbose {
		fmt.Println(opts)
	}

	conn, connErr := db.CreateConnection(opts)
	if connErr != nil {
		log.Fatal("Connection error: " + connErr.Error())
	}
	repo := db.CreateRepo(conn)
	defer repo.CloseConn()

	for _, o := range models.GetTableObjects() {
		if opts.Verbose {
			log.Printf("File Name: %s\n", o.GetFileName())
		}

		csvFile, fileErr := os.Open(o.GetFilePath())
		if fileErr != nil {
			log.Printf("Error: Unable to open file '%s'. %s\n", o.GetFilePath(), fileErr)
			continue
		}

		psFunc, genErr := o.GenParseAndStoreCSV(csvFile, repo, gocsv.UnmarshalFile)
		if genErr != nil {
			log.Println("ERROR: There was an issue generating the ParseAnStoreCSV func")
		} else {
			psErr := psFunc()
			if psErr != nil {
				log.Printf("There was an error while attempting to parse and storethe file %s\nError: %s\n", o.GetFilePath(), psErr.Error())
			}
		}
		fileErr = csvFile.Close()
		if fileErr != nil {
			fmt.Printf("An error occurred while attempting to close the file '%s'. Error: %s\n", o.GetFileName(), fileErr.Error())
		}
	}
}
