package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// SeriesPost is a model that maps the CSV to a DB Table
type SeriesPost struct {
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round,omitempty"  bson:"round"`
   Teamidwinner   string `json:"teamIDwinner"  csv:"teamIDwinner"  db:"teamIDwinner,omitempty"  bson:"teamIDwinner"`
   Lgidwinner   string `json:"lgIDwinner"  csv:"lgIDwinner"  db:"lgIDwinner,omitempty"  bson:"lgIDwinner"`
   Teamidloser   string `json:"teamIDloser"  csv:"teamIDloser"  db:"teamIDloser,omitempty"  bson:"teamIDloser"`
   Lgidloser   string `json:"lgIDloser"  csv:"lgIDloser"  db:"lgIDloser,omitempty"  bson:"lgIDloser"`
   Wins   int `json:"wins"  csv:"wins"  db:"wins"  bson:"wins"`
   Losses   int `json:"losses"  csv:"losses"  db:"losses"  bson:"losses"`
   Ties   int `json:"ties"  csv:"ties"  db:"ties"  bson:"ties"`
  inputDir string
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
  return filepath.Join(m.inputDir, "SeriesPost.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *SeriesPost) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
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
