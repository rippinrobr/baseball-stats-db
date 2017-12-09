package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
	"github.com/rippinrobr/baseball-stats-db/internal/retrosheet"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

const (
	dbpathParam       string = "dbpath"
	dbhostParam              = "dbhost"
	dbnameParam              = "dbname"
	dbpassParam              = "dbpass"
	dbportParam              = "dbport"
	dbtypeParam              = "dbtype"
	dbuserParam              = "dbuser"
	filePattern              = "SKED.TXT"
	inputdirParam            = "inputdir"
	scheduleTableName        = "schedules"
	verboseParam             = "verbose"
)

// main is the entry point for the application.
func main() {
	log.Println("main : Started")

	// Check the environment for a configured port value.
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "retrosheet_schedules.sqlite3"
	}

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
	flag.StringVar(&dbtype, dbtypeParam, "", "indicates what type of database is the load target. Supported databases are MongoDB, Postgres, and SQLite")
	flag.StringVar(&dbuser, dbuserParam, "", "the username to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&inputdir, inputdirParam, "", "the directory where the Retrosheet.org Schedule CSV files live. Required")
	flag.BoolVar(&verbose, verboseParam, false, "writes more lines to the logs")
	flag.Parse()

	dbtype = strings.ToLower(dbtype)
	if !db.IsSupportedDB(dbtype) {
		flag.Usage()
		fmt.Printf("\n'%s' is not a supported database\n", dbtype)
		os.Exit(1)
	}

	if dbtype == db.DBSQLite && dbpath == "" {
		dbpath = "./retrosheet.sqlite3"
		log.Printf("No -dbpath value provided defaulting to '%s'\n", dbpath)
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
	conn, connErr := db.CreateConnection(opts)
	if connErr != nil {
		log.Fatal("Connection error: " + connErr.Error())
	}
	repo = db.CreateRepo(conn)

	defer repo.CloseConn()

	processFiles(repo, opts.Verbose, inputdir)
}

func processFiles(repo db.Repository, isVerbose bool, inDir string) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fname := f.Name()
		if strings.HasSuffix(fname, filePattern) {
			season, err := strconv.Atoi(fname[0:4])
			file, err := os.Open(filepath.Join(inDir, fname))
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			r := csv.NewReader(file)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}

				// if the len(record) == 1 then its the empty line at the end of the file
				// and throws an error because there aren't 12 columns like there should be
				// so I'm going to just continue on
				if err != nil {
					continue
				}

				s, serr := retrosheet.NewSchedule(season, record)
				if serr != nil {
					log.Printf("an error occurred while trying to create schedule object %s\n", serr.Error())
				}

				err = repo.Insert(scheduleTableName, s)
				if err != nil {
					log.Println("error inserting schedule entry", err.Error())
					log.Printf("Schedule\n=========================\n%+v\n", s)
				}
			}
		}
	}

}
