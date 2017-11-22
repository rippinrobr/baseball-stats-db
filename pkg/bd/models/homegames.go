package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-db/pkg/db"

)

// HomeGames is a model that maps the CSV to a DB Table
type HomeGames struct {
   Yearkey   int `json:"yearkey"  csv:"year.key"  db:"yearkey"`
   Leaguekey   string `json:"leaguekey"  csv:"league.key"  db:"leaguekey,omitempty"`
   Teamkey   string `json:"teamkey"  csv:"team.key"  db:"teamkey,omitempty"`
   Parkkey   string `json:"parkkey"  csv:"park.key"  db:"parkkey,omitempty"`
   Spanfirst   string `json:"spanfirst"  csv:"span.first"  db:"spanfirst,omitempty"`
   Spanlast   string `json:"spanlast"  csv:"span.last"  db:"spanlast,omitempty"`
   Games   int `json:"games"  csv:"games"  db:"games"`
   Openings   int `json:"openings"  csv:"openings"  db:"openings"`
   Attendance   int `json:"attendance"  csv:"attendance"  db:"attendance"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *HomeGames) GetTableName() string {
  return "homegames"
}

// GetFileName returns the name of the source file the model was created from
func (m *HomeGames) GetFileName() string {
  return "HomeGames.csv"
}

// GetFilePath returns the path of the source file
func (m *HomeGames) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/HomeGames.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *HomeGames) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*HomeGames, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("HomeGames ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("HomeGames ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("HomeGames ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
