package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Managers is a model that maps the CSV to a DB Table
type Managers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Inseason   int `json:"inseason"  csv:"inseason"  db:"inseason"`
   G   int `json:"g"  csv:"G"  db:"G"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   Rank   string `json:"rank"  csv:"rank"  db:"rank"`
   Plyrmgr   string `json:"plyrmgr"  csv:"plyrMgr"  db:"plyrMgr"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Managers) GetTableName() string {
  return "managers"
}

// GetFileName returns the name of the source file the model was created from
func (m *Managers) GetFileName() string {
  return "Managers.csv"
}

// GetFilePath returns the path of the source file
func (m *Managers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Managers.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Managers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*Managers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Managers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Managers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Managers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
