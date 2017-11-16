package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// AllstarFull is a model that maps the CSV to a DB Table
type AllstarFull struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Gamenum   int `json:"gamenum"  csv:"gameNum"  db:"gameNum"`
   Gameid   string `json:"gameid"  csv:"gameID"  db:"gameID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Gp   string `json:"gp"  csv:"GP"  db:"GP"`
   Startingpos   string `json:"startingpos"  csv:"startingPos"  db:"startingPos"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AllstarFull) GetTableName() string {
  return "allstarfull"
}

// GetFileName returns the name of the source file the model was created from
func (m *AllstarFull) GetFileName() string {
  return "AllstarFull.csv"
}

// GetFilePath returns the path of the source file
func (m *AllstarFull) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AllstarFull.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AllstarFull) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*AllstarFull, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AllstarFull ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AllstarFull ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AllstarFull ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
