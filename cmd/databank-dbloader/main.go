package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"

	"github.com/rippinrobr/baseball-stats-db/internal/databank"
	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
)

const (
	dbpathParam   string = "dbpath"
	dbhostParam          = "dbhost"
	dbnameParam          = "dbname"
	dbpassParam          = "dbpass"
	dbportParam          = "dbport"
	dbtypeParam          = "dbtype"
	dbuserParam          = "dbuser"
	inputdirParam        = "inputdir"
	verboseParam         = "verbose"
)

func main() {
	// cmd line args
	var dbhost, dbname, dbpass, dbpath, dbtype, dbuser, inputdir string
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
	flag.StringVar(&inputdir, inputdirParam, "", "the directory where the Baseball Databank CSV files live. Required")
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

	if inputdir == "" {
		flag.Usage()
		fmt.Printf("\n-inputdir is required.")
		os.Exit(1)
	}

	stat, inputdirErr := os.Stat(inputdir)
	if inputdirErr != nil {
		log.Fatalf("The directory '%s' does not exist. Please provide a valid directory for -inputdir", inputdir)
	}

	if !stat.IsDir() {
		log.Fatalf("The path '%s' is not a directory. Please provide the path for the directory that contains the CSV files", inputdir)
	}

	opts := db.Options{
		Host:    dbhost,
		Name:    dbname,
		Pass:    dbpass,
		Path:    dbpath,
		Port:    dbport,
		Type:    dbtype,
		User:    dbuser,
		Verbose: verbose,
	}

	if opts.Verbose {
		fmt.Println(opts)
	}

	var repo db.Repository

	switch dbtype {
	case db.DBMongoDB:
		conn, err := db.CreateNoSQLConnection(opts)
		if err != nil {
			log.Fatal("Connection error: " + err.Error())
		}
		repo = db.CreateNoSQLRepo(conn)
	default:
		conn, connErr := db.CreateConnection(opts)
		if connErr != nil {
			log.Fatal("Connection error: " + connErr.Error())
		}
		repo = db.CreateRepo(conn)
	}
	defer repo.CloseConn()

	processFiles(repo, opts.Verbose, inputdir)
}

func processFiles(repo db.Repository, isVerbose bool, inDir string) {
	for _, o := range databank.GetTableObjects() {
		if isVerbose {
			log.Printf("File Name: %s\n", o.GetFileName())
		}

		o.SetInputDirectory(inDir)
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
