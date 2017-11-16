package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Salaries is a model that maps the CSV to a DB Table
type Salaries struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
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
func (m *Salaries) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
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
   }
}
