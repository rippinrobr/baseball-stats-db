package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// CollegePlaying is a model that maps the CSV to a DB Table
type CollegePlaying struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Schoolid   string `json:"schoolid"  csv:"schoolID"  db:"schoolID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *CollegePlaying) GetTableName() string {
  return "collegeplaying"
}

// GetFileName returns the name of the source file the model was created from
func (m *CollegePlaying) GetFileName() string {
  return "CollegePlaying.csv"
}

// GetFilePath returns the path of the source file
func (m *CollegePlaying) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/CollegePlaying.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *CollegePlaying) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*CollegePlaying, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("CollegePlaying ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("CollegePlaying ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("CollegePlaying ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
