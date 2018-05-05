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
	dbpathParam     string = "dbpath"
	dbhostParam            = "dbhost"
	dbnameParam            = "dbname"
	dbpassParam            = "dbpass"
	dbportParam            = "dbport"
	dbtypeParam            = "dbtype"
	dbuserParam            = "dbuser"
	inputdirParam          = "inputdir"
	inputfilesParam        = "inputfiles"
	verboseParam           = "verbose"
)

func main() {
	// cmd line args
	var dbhost, dbname, dbpass, dbpath, dbtype, dbuser, inputdir, inputfiles string
	var dbport int
	var verbose bool
	// since SQLite is the only supported db at the moment then I will default to it for now
	flag.StringVar(&dbhost, dbhostParam, "localhost", "the name or ip address of the database server.")
	flag.StringVar(&dbname, dbnameParam, "", "the hame of the database to load")
	flag.StringVar(&dbpass, dbpassParam, "", "the password to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbpath, dbpathParam, "", "the path to your SQLite database")
	flag.IntVar(&dbport, dbportParam, 0, "the port to use when connecting to the database the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbtype, dbtypeParam, "", "indicates what type of database is the load target. Supported databases are MongoDB, Postgres, and SQLite")
	flag.StringVar(&dbuser, dbuserParam, "", "the username to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&inputdir, inputdirParam, "", "the directory where the Baseball Databank CSV files live. Required")
	flag.StringVar(&inputfiles, inputfilesParam, "", "a comma-demilited string of file names to parse and load.")
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

	if inputdir == "" && inputfiles == "" {
		flag.Usage()
		fmt.Printf("\n-inputdir or inputfile is required.")
		os.Exit(1)
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

	// if the inputdir value is set bu the inputfiles value is not then
	// I'm processing all of the files
	if inputdir != "" && inputfiles == "" {
		stat, inputdirErr := os.Stat(inputdir)
		if inputdirErr != nil {
			log.Fatalf("The directory '%s' does not exist. Please provide a valid directory for -inputdir", inputdir)
		}

		if !stat.IsDir() {
			log.Fatalf("The path '%s' is not a directory. Please provide the path for the directory that contains the CSV files", inputdir)
		}

		processInputDirectory(repo, opts.Verbose, inputdir)
		return
	}

	if inputdir == "" {
		fmt.Println("inputdir must be set to the base directory when using the -inputfiles option")
		os.Exit(1)
	}
	processInputFiles(repo, opts.Verbose, inputdir, strings.Split(inputfiles, ","))
}

func processInputFiles(repo db.Repository, isVerbose bool, inputdir string, fileNames []string) {
	for _, f := range fileNames {
		o := databank.LookupTableObject(f)
		if o == nil {
			fmt.Printf("ERROR: Unable to find a match for the file '%s'\n", f)
			continue
		}
		o.SetInputDirectory(inputdir)
		parseAndStoreFile(o, repo)
	}
}

func processInputDirectory(repo db.Repository, isVerbose bool, inDir string) {
	for _, o := range databank.GetTableObjects() {
		if isVerbose {
			log.Printf("File Name: %s\n", o.GetFileName())
		}

		o.SetInputDirectory(inDir)
		// csvFile, fileErr := os.Open(o.GetFilePath())
		// if fileErr != nil {
		// 	log.Printf("Error: Unable to open file '%s'. %s\n", o.GetFilePath(), fileErr)
		// 	continue
		// }

		parseAndStoreFile(o, repo)
	}
}

func parseAndStoreFile(o databank.TableObject, repo db.Repository) {

	csvFile, fileErr := os.Open(o.GetFilePath())
	if fileErr != nil {
		log.Printf("Error: Unable to open file '%s'. %s\n", o.GetFilePath(), fileErr)
		return
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
