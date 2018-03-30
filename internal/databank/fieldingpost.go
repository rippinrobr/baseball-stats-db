package databank

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
)

// FieldingPost is a model that maps the CSV to a DB Table
type FieldingPost struct {
	Playerid string  `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
	Yearid   int     `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
	Teamid   string  `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"  bson:"teamID"`
	Lgid     string  `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
	Round    string  `json:"round"  csv:"round"  db:"round,omitempty"  bson:"round"`
	Pos      string  `json:"pOS"  csv:"POS"  db:"POS,omitempty"  bson:"POS"`
	G        int     `json:"g"  csv:"G"  db:"G"  bson:"G"`
	Gs       int     `json:"gS"  csv:"GS"  db:"GS"  bson:"GS"`
	Innouts  float32 `json:"innOuts"  csv:"InnOuts"  db:"InnOuts"  bson:"InnOuts"`
	Po       int     `json:"pO"  csv:"PO"  db:"PO"  bson:"PO"`
	A        int     `json:"a"  csv:"A"  db:"A"  bson:"A"`
	E        int     `json:"e"  csv:"E"  db:"E"  bson:"E"`
	Dp       int     `json:"dP"  csv:"DP"  db:"DP"  bson:"DP"`
	Tp       float32 `json:"tP"  csv:"TP"  db:"TP"  bson:"TP"`
	Pb       int     `json:"pB"  csv:"PB"  db:"PB"  bson:"PB"`
	Sb       int     `json:"sB"  csv:"SB"  db:"SB"  bson:"SB"`
	Cs       int     `json:"cS"  csv:"CS"  db:"CS"  bson:"CS"`
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
