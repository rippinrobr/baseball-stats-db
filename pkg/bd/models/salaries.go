package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-db/pkg/db"

)

// Salaries is a model that maps the CSV to a DB Table
type Salaries struct {
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
   Salary   int `json:"salary"  csv:"salary"  db:"salary"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Salaries) GetTableName() string {
  return "salaries"
}

// GetFileName returns the name of the source file the model was created from
func (m *Salaries) GetFileName() string {
  return "Salaries.csv"
}

// GetFilePath returns the path of the source file
func (m *Salaries) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Salaries.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Salaries) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Salaries, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Salaries ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Salaries ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Salaries ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
