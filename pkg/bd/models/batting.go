package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Batting is a model that maps the CSV to a DB Table
type Batting struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ab   int `json:"ab"  csv:"AB"  db:"AB"`
   R   int `json:"r"  csv:"R"  db:"R"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Doubles   int `json:"doubles"  csv:"2B"  db:"2B"`
   Triples   int `json:"triples"  csv:"3B"  db:"3B"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Rbi   string `json:"rbi"  csv:"RBI"  db:"RBI"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   string `json:"so"  csv:"SO"  db:"SO"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Batting) GetTableName() string {
  return "batting"
}

// GetFileName returns the name of the source file the model was created from
func (m *Batting) GetFileName() string {
  return "Batting.csv"
}

// GetFilePath returns the path of the source file
func (m *Batting) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Batting.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Batting) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Batting, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Batting ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Batting ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Batting ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
