package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// AwardsSharePlayers is a model that maps the CSV to a DB Table
type AwardsSharePlayers struct {
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Pointswon   string `json:"pointswon"  csv:"pointsWon"  db:"pointsWon"`
   Pointsmax   int `json:"pointsmax"  csv:"pointsMax"  db:"pointsMax"`
   Votesfirst   string `json:"votesfirst"  csv:"votesFirst"  db:"votesFirst"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AwardsSharePlayers) GetTableName() string {
  return "awardsshareplayers"
}

// GetFileName returns the name of the source file the model was created from
func (m *AwardsSharePlayers) GetFileName() string {
  return "AwardsSharePlayers.csv"
}

// GetFilePath returns the path of the source file
func (m *AwardsSharePlayers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsSharePlayers.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AwardsSharePlayers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*AwardsSharePlayers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AwardsSharePlayers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AwardsSharePlayers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AwardsSharePlayers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
