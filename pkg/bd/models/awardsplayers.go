package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// AwardsPlayers is a model that maps the CSV to a DB Table
type AwardsPlayers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Tie   string `json:"tie"  csv:"tie"  db:"tie"`
   Notes   string `json:"notes"  csv:"notes"  db:"notes"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AwardsPlayers) GetTableName() string {
  return "awardsplayers"
}

// GetFileName returns the name of the source file the model was created from
func (m *AwardsPlayers) GetFileName() string {
  return "AwardsPlayers.csv"
}

// GetFilePath returns the path of the source file
func (m *AwardsPlayers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsPlayers.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AwardsPlayers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*AwardsPlayers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AwardsPlayers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AwardsPlayers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AwardsPlayers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
