package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// HallOfFame is a model that maps the CSV to a DB Table
type HallOfFame struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearid"  db:"yearid"`
   Votedby   string `json:"votedby"  csv:"votedBy"  db:"votedBy"`
   Ballots   string `json:"ballots"  csv:"ballots"  db:"ballots"`
   Needed   string `json:"needed"  csv:"needed"  db:"needed"`
   Votes   string `json:"votes"  csv:"votes"  db:"votes"`
   Inducted   string `json:"inducted"  csv:"inducted"  db:"inducted"`
   Category   string `json:"category"  csv:"category"  db:"category"`
   Needednote   string `json:"needednote"  csv:"needed_note"  db:"needed_note"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *HallOfFame) GetTableName() string {
  return "halloffame"
}

// GetFileName returns the name of the source file the model was created from
func (m *HallOfFame) GetFileName() string {
  return "HallOfFame.csv"
}

// GetFilePath returns the path of the source file
func (m *HallOfFame) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/HallOfFame.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *HallOfFame) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*HallOfFame, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("HallOfFame ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("HallOfFame ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("HallOfFame ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
