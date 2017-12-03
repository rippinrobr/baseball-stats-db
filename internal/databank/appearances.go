package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// Appearances is a model that maps the CSV to a DB Table
type Appearances struct {
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"  bson:"teamID"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
   Gall   int `json:"gall"  csv:"G_all"  db:"G_all"  bson:"G_all"`
   Gs   int `json:"gS"  csv:"GS"  db:"GS"  bson:"GS"`
   Gbatting   int `json:"gbatting"  csv:"G_batting"  db:"G_batting"  bson:"G_batting"`
   Gdefense   int `json:"gdefense"  csv:"G_defense"  db:"G_defense"  bson:"G_defense"`
   Gp   int `json:"gp"  csv:"G_p"  db:"G_p"  bson:"G_p"`
   Gc   int `json:"gc"  csv:"G_c"  db:"G_c"  bson:"G_c"`
   G1B   int `json:"g1b"  csv:"G_1b"  db:"G_1b"  bson:"G_1b"`
   G2B   int `json:"g2b"  csv:"G_2b"  db:"G_2b"  bson:"G_2b"`
   G3B   int `json:"g3b"  csv:"G_3b"  db:"G_3b"  bson:"G_3b"`
   Gss   int `json:"gss"  csv:"G_ss"  db:"G_ss"  bson:"G_ss"`
   Glf   int `json:"glf"  csv:"G_lf"  db:"G_lf"  bson:"G_lf"`
   Gcf   int `json:"gcf"  csv:"G_cf"  db:"G_cf"  bson:"G_cf"`
   Grf   int `json:"grf"  csv:"G_rf"  db:"G_rf"  bson:"G_rf"`
   Gof   int `json:"gof"  csv:"G_of"  db:"G_of"  bson:"G_of"`
   Gdh   int `json:"gdh"  csv:"G_dh"  db:"G_dh"  bson:"G_dh"`
   Gph   int `json:"gph"  csv:"G_ph"  db:"G_ph"  bson:"G_ph"`
   Gpr   int `json:"gpr"  csv:"G_pr"  db:"G_pr"  bson:"G_pr"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Appearances) GetTableName() string {
  return "appearances"
}

// GetFileName returns the name of the source file the model was created from
func (m *Appearances) GetFileName() string {
  return "Appearances.csv"
}

// GetFilePath returns the path of the source file
func (m *Appearances) GetFilePath() string {
  return filepath.Join(m.inputDir, "Appearances.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *Appearances) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Appearances) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Appearances, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Appearances ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Appearances ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Appearances ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
