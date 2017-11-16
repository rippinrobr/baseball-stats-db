package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// AwardsManagers is a model that maps the CSV to a DB Table
type AwardsManagers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Tie   string `json:"tie"  csv:"tie"  db:"tie"`
   Notes   string `json:"notes"  csv:"notes"  db:"notes"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AwardsManagers) GetTableName() string {
  return "awardsmanagers"
}

// GetFileName returns the name of the source file the model was created from
func (m *AwardsManagers) GetFileName() string {
  return "AwardsManagers.csv"
}

// GetFilePath returns the path of the source file
func (m *AwardsManagers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsManagers.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AwardsManagers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*AwardsManagers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AwardsManagers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AwardsManagers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AwardsManagers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
