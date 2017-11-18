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

	"github.com/gocarina/gocsv"

	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"
	"github.com/rippinrobr/baseball-databank-tools/pkg/db"
)

type optionalConfig struct {
	Verbose bool
}

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
		log.Fatalf("A -dbconn value is required for the database type '%s'\n", dbtype)
	}

	processFiles(dbtype, dbconn, &optionalConfig{verbose})

}

func processFiles(dbtype, dbconn string, opts *optionalConfig) {
	conn, connErr := db.CreateConnection(dbtype, dbconn)
	if connErr != nil {
		log.Fatal("Connection error: " + connErr.Error())
	}
	repo := db.CreateRepo(conn)
	defer repo.CloseConn()

	for _, o := range models.GetTableObjects() {
		if opts.Verbose {
			//log.Printf("File Path: '%s'\n", o.GetFilePath())
			log.Printf("File Name: %s\n", o.GetFileName())
		}

		csvFile, fileErr := os.Open(o.GetFilePath())
		if fileErr != nil {
			log.Printf("Error: Unable to open file '%s'. %s\n", o.GetFilePath(), fileErr)
			continue
		}

		psFunc, genErr := o.GenParseAndStoreCSV(csvFile, repo, gocsv.UnmarshalFile)
		if genErr != nil {
			log.Println("ERROR: There was an issue generating the ParseAnStoreCSV func\n")
		} else {
			psErr := psFunc()
			if psErr != nil {
				log.Printf("There was an error while attempting to parse and storethe file %s\nError: %s\n", o.GetFilePath(), psErr.Error())
			}
		}
		csvFile.Close()
		// if dbtype == db.DBSQLite {
		// 	log.Println("Since I'm loading SQLite I need to sleep 10 secs so SQLite can catch up")
		// 	time.Sleep(10 * time.Second)
		// 	log.Println("I'm awake")
		// }
	}
}
