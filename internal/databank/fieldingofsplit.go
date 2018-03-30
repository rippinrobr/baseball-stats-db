package databank

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
)

// FieldingOFsplit is a model that maps the CSV to a DB Table
type FieldingOFsplit struct {
	Playerid string  `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
	Yearid   int     `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
	Stint    int     `json:"stint"  csv:"stint"  db:"stint"  bson:"stint"`
	Teamid   string  `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"  bson:"teamID"`
	Lgid     string  `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
	Pos      string  `json:"pOS"  csv:"POS"  db:"POS,omitempty"  bson:"POS"`
	G        int     `json:"g"  csv:"G"  db:"G"  bson:"G"`
	Gs       int     `json:"gS"  csv:"GS"  db:"GS"  bson:"GS"`
	Innouts  int     `json:"innOuts"  csv:"InnOuts"  db:"InnOuts"  bson:"InnOuts"`
	Po       float32 `json:"pO"  csv:"PO"  db:"PO"  bson:"PO"`
	A        float32 `json:"a"  csv:"A"  db:"A"  bson:"A"`
	E        float32 `json:"e"  csv:"E"  db:"E"  bson:"E"`
	Dp       float32 `json:"dP"  csv:"DP"  db:"DP"  bson:"DP"`
	Pb       string  `json:"pB"  csv:"PB"  db:"PB,omitempty"  bson:"PB"`
	Wp       string  `json:"wP"  csv:"WP"  db:"WP,omitempty"  bson:"WP"`
	Sb       string  `json:"sB"  csv:"SB"  db:"SB,omitempty"  bson:"SB"`
	Cs       string  `json:"cS"  csv:"CS"  db:"CS,omitempty"  bson:"CS"`
	Zr       string  `json:"zR"  csv:"ZR"  db:"ZR,omitempty"  bson:"ZR"`
	inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *FieldingOFsplit) GetTableName() string {
	return "fieldingofsplit"
}

// GetFileName returns the name of the source file the model was created from
func (m *FieldingOFsplit) GetFileName() string {
	return "FieldingOFsplit.csv"
}

// GetFilePath returns the path of the source file
func (m *FieldingOFsplit) GetFilePath() string {
	return filepath.Join(m.inputDir, "FieldingOFsplit.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *FieldingOFsplit) SetInputDirectory(inputDir string) {
	m.inputDir = inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingOFsplit) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
	if f == nil {
		return func() error { return nil }, errors.New("nil File")
	}

	return func() error {
		rows := make([]*FieldingOFsplit, 0)
		numErrors := 0
		err := pfunc(f, &rows)
		if err == nil {
			numrows := len(rows)
			if numrows > 0 {
				log.Println("FieldingOFsplit ==> Truncating")
				terr := repo.Truncate(m.GetTableName())
				if terr != nil {
					log.Println("truncate err:", terr.Error())
				}

				log.Printf("FieldingOFsplit ==> Inserting %d Records\n", numrows)
				for _, r := range rows {
					ierr := repo.Insert(m.GetTableName(), r)
					if ierr != nil {
						log.Printf("Insert error: %s\n", ierr.Error())
						numErrors++
					}
				}
			}
			log.Printf("FieldingOFsplit ==> %d records created\n", numrows-numErrors)
		}

		return err
	}, nil
}
