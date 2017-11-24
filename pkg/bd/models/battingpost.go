package models


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-db/pkg/db"

)

// BattingPost is a model that maps the CSV to a DB Table
type BattingPost struct {
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round,omitempty"`
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ab   int `json:"aB"  csv:"AB"  db:"AB"`
   R   int `json:"r"  csv:"R"  db:"R"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Doubles   int `json:"doubles"  csv:"2B"  db:"doubles"`
   Triples   int `json:"triples"  csv:"3B"  db:"triples"`
   Hr   int `json:"hR"  csv:"HR"  db:"HR"`
   Rbi   int `json:"rBI"  csv:"RBI"  db:"RBI"`
   Sb   int `json:"sB"  csv:"SB"  db:"SB"`
   Cs   int `json:"cS"  csv:"CS"  db:"CS"`
   Bb   int `json:"bB"  csv:"BB"  db:"BB"`
   So   int `json:"sO"  csv:"SO"  db:"SO"`
   Ibb   int `json:"iBB"  csv:"IBB"  db:"IBB"`
   Hbp   int `json:"hBP"  csv:"HBP"  db:"HBP"`
   Sh   int `json:"sH"  csv:"SH"  db:"SH"`
   Sf   int `json:"sF"  csv:"SF"  db:"SF"`
   Gidp   int `json:"gIDP"  csv:"GIDP"  db:"GIDP"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *BattingPost) GetTableName() string {
  return "battingpost"
}

// GetFileName returns the name of the source file the model was created from
func (m *BattingPost) GetFileName() string {
  return "BattingPost.csv"
}

// GetFilePath returns the path of the source file
func (m *BattingPost) GetFilePath() string {
  return filepath.Join(m.inputDir, "BattingPost.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *BattingPost) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *BattingPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*BattingPost, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("BattingPost ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("BattingPost ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("BattingPost ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
