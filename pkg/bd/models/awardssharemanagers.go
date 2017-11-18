package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// AwardsShareManagers is a model that maps the CSV to a DB Table
type AwardsShareManagers struct {
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Pointswon   int `json:"pointswon"  csv:"pointsWon"  db:"pointsWon"`
   Pointsmax   string `json:"pointsmax"  csv:"pointsMax"  db:"pointsMax"`
   Votesfirst   int `json:"votesfirst"  csv:"votesFirst"  db:"votesFirst"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AwardsShareManagers) GetTableName() string {
  return "awardssharemanagers"
}

// GetFileName returns the name of the source file the model was created from
func (m *AwardsShareManagers) GetFileName() string {
  return "AwardsShareManagers.csv"
}

// GetFilePath returns the path of the source file
func (m *AwardsShareManagers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsShareManagers.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AwardsShareManagers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*AwardsShareManagers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AwardsShareManagers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AwardsShareManagers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AwardsShareManagers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
