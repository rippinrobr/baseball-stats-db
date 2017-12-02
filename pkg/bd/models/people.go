package models

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/rippinrobr/baseball-stats-db/pkg/parsers/csv"

	"github.com/rippinrobr/baseball-stats-db/pkg/db"
)

// People is a model that maps the CSV to a DB Table
type People struct {
	Playerid     string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
	Birthyear    int    `json:"birthYear"  csv:"birthYear"  db:"birthYear"`
	Birthmonth   int    `json:"birthMonth"  csv:"birthMonth"  db:"birthMonth"`
	Birthday     int    `json:"birthDay"  csv:"birthDay"  db:"birthDay"`
	Birthcountry string `json:"birthCountry"  csv:"birthCountry"  db:"birthCountry,omitempty"`
	Birthstate   string `json:"birthState"  csv:"birthState"  db:"birthState,omitempty"`
	Birthcity    string `json:"birthCity"  csv:"birthCity"  db:"birthCity,omitempty"`
	Deathyear    int    `json:"deathYear"  csv:"deathYear"  db:"deathYear"`
	Deathmonth   int    `json:"deathMonth"  csv:"deathMonth"  db:"deathMonth"`
	Deathday     int    `json:"deathDay"  csv:"deathDay"  db:"deathDay"`
	Deathcountry string `json:"deathCountry"  csv:"deathCountry"  db:"deathCountry,omitempty"`
	Deathstate   string `json:"deathState"  csv:"deathState"  db:"deathState,omitempty"`
	Deathcity    string `json:"deathCity"  csv:"deathCity"  db:"deathCity,omitempty"`
	Namefirst    string `json:"nameFirst"  csv:"nameFirst"  db:"nameFirst,omitempty"`
	Namelast     string `json:"nameLast"  csv:"nameLast"  db:"nameLast,omitempty"`
	Namegiven    string `json:"nameGiven"  csv:"nameGiven"  db:"nameGiven,omitempty"`
	Weight       int    `json:"weight"  csv:"weight"  db:"weight"`
	Height       int    `json:"height"  csv:"height"  db:"height"`
	Bats         string `json:"bats"  csv:"bats"  db:"bats,omitempty"`
	Throws       string `json:"throws"  csv:"throws"  db:"throws,omitempty"`
	Debut        string `json:"debut"  csv:"debut"  db:"debut,omitempty"`
	Finalgame    string `json:"finalGame"  csv:"finalGame"  db:"finalGame,omitempty"`
	Retroid      string `json:"retroID"  csv:"retroID"  db:"retroID,omitempty"`
	Bbrefid      string `json:"bbrefID"  csv:"bbrefID"  db:"bbrefID,omitempty"`
	inputDir     string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *People) GetTableName() string {
	return "people"
}

// GetFileName returns the name of the source file the model was created from
func (m *People) GetFileName() string {
	return "People.csv"
}

// GetFilePath returns the path of the source file
func (m *People) GetFilePath() string {
	return filepath.Join(m.inputDir, "People.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *People) SetInputDirectory(inputDir string) {
	m.inputDir = inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *People) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
	if f == nil {
		return func() error { return nil }, errors.New("nil File")
	}

	return func() error {
		rows := make([]*People, 0)
		numErrors := 0
		err := pfunc(f, &rows)
		if err == nil {
			numrows := len(rows)
			if numrows > 0 {
				log.Println("People ==> Truncating")
				terr := repo.Truncate(m.GetTableName())
				if terr != nil {
					log.Println("truncate err:", terr.Error())
				}

				log.Printf("People ==> Inserting %d Records\n", numrows)
				for _, r := range rows {
					ierr := repo.Insert(m.GetTableName(), r)
					if ierr != nil {
						log.Printf("Insert error: %s\n", ierr.Error())
						numErrors++
					}
				}
			}
			log.Printf("People ==> %d records created\n", numrows-numErrors)
		}

		return err
	}, nil
}
