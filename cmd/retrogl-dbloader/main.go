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
	dbpathParam      string = "dbpath"
	dbhostParam             = "dbhost"
	dbnameParam             = "dbname"
	dbpassParam             = "dbpass"
	dbportParam             = "dbport"
	dbtypeParam             = "dbtype"
	dbuserParam             = "dbuser"
	filePrefix              = "GL"
	fileSuffix              = ".TXT"
	inputdirParam           = "inputdir"
	inputFileParam          = "inputfile"
	gamelogTableName        = "gamelogs"
	verboseParam            = "verbose"
)

// main is the entry point for the application.
func main() {
	log.Println("main : Started")

	// cmd line args
	var dbhost, dbname, dbpass, dbpath, dbtype, dbuser, inputdir, inputfile string
	var dbport int
	var verbose bool

	// since SQLite is the only supported db at the moment then I will default to it for now
	flag.StringVar(&dbhost, dbhostParam, "localhost", "the name or ip address of the database server.")
	flag.StringVar(&dbname, dbnameParam, "", "the hame of the database to load")
	flag.StringVar(&dbpass, dbpassParam, "", "the password to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbpath, dbpathParam, "", "the path to your SQLite database")
	flag.IntVar(&dbport, dbportParam, 0, "the port to use when connecting to the database the database. Required for all dbtypes except SQLite")
	flag.StringVar(&dbtype, dbtypeParam, "", "indicates what type of database is the load target. Supported databases are MongoDB, MySQL, Postgres, and SQLite")
	flag.StringVar(&dbuser, dbuserParam, "", "the username to use when loading the database. Required for all dbtypes except SQLite")
	flag.StringVar(&inputdir, inputdirParam, "", "the directory where the Retrosheet.org Game Log CSV files live. One of the following is required: -inputfile or -inputdir.")
	flag.StringVar(&inputfile, inputFileParam, "", "a Retrosheet.org Game Log CSV file to be parsed and stored. One of the following is required: -inputfile or -inputdir.")
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

	if inputdir == "" && inputfile == "" {
		flag.Usage()
		fmt.Printf("\neither -inputdir or -inputfile is required.")
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

	if inputdir != "" {
		stat, inputdirErr := os.Stat(inputdir)
		if inputdirErr != nil {
			log.Fatalf("The directory '%s' does not exist. Please provide a valid directory for -inputdir", inputdir)
		}

		if !stat.IsDir() {
			log.Fatalf("The path '%s' is not a directory. Please provide the path for the directory that contains the CSV files", inputdir)
		}

		processDirectory(repo, opts.Verbose, inputdir)
		return
	}

	if _, err := os.Stat(inputfile); err != nil {
		log.Fatalf("The file '%s' does not exist. Please provide a valid Retrosheet.org game log CSV file for -inputfile", inputfile)
	}

	processFile(repo, opts.Verbose, inputfile)
}

func processDirectory(repo db.Repository, isVerbose bool, inDir string) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	repo.Truncate(gamelogTableName)
	for _, f := range files {
		fname := f.Name()
		if strings.HasPrefix(fname, filePrefix) && strings.HasSuffix(fname, fileSuffix) {
			season, err := strconv.Atoi(fname[2:6])
			if err != nil {
				fmt.Println("err", err.Error())
			}
			fmt.Println("season: ", season)
			if season == 0 {
				fmt.Println("season was zero so I must be at the end of the file. I'm out!")
				return
			}
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

				gl, glerr := retrosheet.NewGameLog(season, record)
				if glerr != nil {
					log.Printf("an error occurred while trying to create game log object %s\n", glerr.Error())
				}

				err = repo.Insert(gamelogTableName, gl)
				if err != nil {
					log.Println("error inserting gamelog entry", err.Error())
					log.Printf("GameLog\n=========================\n%+v\n", gl)
				}
			}
		}
	}

}

func processFile(repo db.Repository, isVerbose bool, inFile string) {
	_, fname := filepath.Split(inFile)
	if strings.HasPrefix(fname, filePrefix) && strings.HasSuffix(fname, fileSuffix) {
		season, err := strconv.Atoi(fname[2:6])
		if err != nil {
			fmt.Println("err", err.Error())
		}
		fmt.Println("season: ", season)
		file, err := os.Open(inFile)
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
			fmt.Println("len(record): ", len(record))

			gl, glerr := retrosheet.NewGameLog(season, record)
			if glerr != nil {
				log.Printf("an error occurred while trying to create game log object %s\n", glerr.Error())
			}

			err = repo.Insert(gamelogTableName, gl)
			if err != nil {
				log.Println("error inserting gamelog entry", err.Error())
				log.Printf("GameLog\n=========================\n%+v\n", gl)
			}
		}
	}
}
