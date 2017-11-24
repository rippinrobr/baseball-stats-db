package models

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

	"github.com/rippinrobr/baseball-databank-db/pkg/db"
)

// FieldingPost is a model that maps the CSV to a DB Table
type FieldingPost struct {
	Playerid string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
	Yearid   int    `json:"yearID"  csv:"yearID"  db:"yearID"`
	Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
	Lgid     string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
	Round    string `json:"round"  csv:"round"  db:"round,omitempty"`
	Pos      string `json:"pOS"  csv:"POS"  db:"POS,omitempty"`
	G        int    `json:"g"  csv:"G"  db:"G"`
	Gs       int    `json:"gS"  csv:"GS"  db:"GS"`
	Innouts  int    `json:"innOuts"  csv:"InnOuts"  db:"InnOuts"`
	Po       int    `json:"pO"  csv:"PO"  db:"PO"`
	A        int    `json:"a"  csv:"A"  db:"A"`
	E        int    `json:"e"  csv:"E"  db:"E"`
	Dp       int    `json:"dP"  csv:"DP"  db:"DP"`
	Tp       int    `json:"tP"  csv:"TP"  db:"TP"`
	Pb       int    `json:"pB"  csv:"PB"  db:"PB"`
	Sb       int    `json:"sB"  csv:"SB"  db:"SB"`
	Cs       int    `json:"cS"  csv:"CS"  db:"CS"`
	inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *FieldingPost) GetTableName() string {
	return "fieldingpost"
}

// GetFileName returns the name of the source file the model was created from
func (m *FieldingPost) GetFileName() string {
	return "FieldingPost.csv"
}

// GetFilePath returns the path of the source file
func (m *FieldingPost) GetFilePath() string {
	return filepath.Join(m.inputDir, "FieldingPost.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *FieldingPost) SetInputDirectory(inputDir string) {
	m.inputDir = inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
	if f == nil {
		return func() error { return nil }, errors.New("nil File")
	}

	return func() error {
		rows := make([]*FieldingPost, 0)
		numErrors := 0
		err := pfunc(f, &rows)
		if err == nil {
			numrows := len(rows)
			if numrows > 0 {
				log.Println("FieldingPost ==> Truncating")
				terr := repo.Truncate(m.GetTableName())
				if terr != nil {
					log.Println("truncate err:", terr.Error())
				}

				log.Printf("FieldingPost ==> Inserting %d Records\n", numrows)
				for _, r := range rows {
					ierr := repo.Insert(m.GetTableName(), r)
					if ierr != nil {
						log.Printf("Insert error: %s\n", ierr.Error())
						numErrors++
					}
				}
			}
			log.Printf("FieldingPost ==> %d records created\n", numrows-numErrors)
		}

		return err
	}, nil
}
