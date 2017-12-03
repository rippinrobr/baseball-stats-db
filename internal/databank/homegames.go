package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// HomeGames is a model that maps the CSV to a DB Table
type HomeGames struct {
   Yearkey   int `json:"yearkey"  csv:"year.key"  db:"yearkey"  bson:"yearkey"`
   Leaguekey   string `json:"leaguekey"  csv:"league.key"  db:"leaguekey,omitempty"  bson:"leaguekey"`
   Teamkey   string `json:"teamkey"  csv:"team.key"  db:"teamkey,omitempty"  bson:"teamkey"`
   Parkkey   string `json:"parkkey"  csv:"park.key"  db:"parkkey,omitempty"  bson:"parkkey"`
   Spanfirst   string `json:"spanfirst"  csv:"span.first"  db:"spanfirst,omitempty"  bson:"spanfirst"`
   Spanlast   string `json:"spanlast"  csv:"span.last"  db:"spanlast,omitempty"  bson:"spanlast"`
   Games   int `json:"games"  csv:"games"  db:"games"  bson:"games"`
   Openings   int `json:"openings"  csv:"openings"  db:"openings"  bson:"openings"`
   Attendance   int `json:"attendance"  csv:"attendance"  db:"attendance"  bson:"attendance"`
  inputDir string
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
  return filepath.Join(m.inputDir, "HomeGames.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *HomeGames) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
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
