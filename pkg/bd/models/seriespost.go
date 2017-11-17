package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// SeriesPost is a model that maps the CSV to a DB Table
type SeriesPost struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Teamidwinner   string `json:"teamidwinner"  csv:"teamIDwinner"  db:"teamIDwinner"`
   Lgidwinner   string `json:"lgidwinner"  csv:"lgIDwinner"  db:"lgIDwinner"`
   Teamidloser   string `json:"teamidloser"  csv:"teamIDloser"  db:"teamIDloser"`
   Lgidloser   string `json:"lgidloser"  csv:"lgIDloser"  db:"lgIDloser"`
   Wins   int `json:"wins"  csv:"wins"  db:"wins"`
   Losses   int `json:"losses"  csv:"losses"  db:"losses"`
   Ties   int `json:"ties"  csv:"ties"  db:"ties"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *SeriesPost) GetTableName() string {
  return "seriespost"
}

// GetFileName returns the name of the source file the model was created from
func (m *SeriesPost) GetFileName() string {
  return "SeriesPost.csv"
}

// GetFilePath returns the path of the source file
func (m *SeriesPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/SeriesPost.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *SeriesPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*SeriesPost, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("SeriesPost ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("SeriesPost ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("SeriesPost ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
