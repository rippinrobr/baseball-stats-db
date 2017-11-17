package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// People is a model that maps the CSV to a DB Table
type People struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Birthyear   string `json:"birthyear"  csv:"birthYear"  db:"birthYear"`
   Birthmonth   string `json:"birthmonth"  csv:"birthMonth"  db:"birthMonth"`
   Birthday   string `json:"birthday"  csv:"birthDay"  db:"birthDay"`
   Birthcountry   string `json:"birthcountry"  csv:"birthCountry"  db:"birthCountry"`
   Birthstate   string `json:"birthstate"  csv:"birthState"  db:"birthState"`
   Birthcity   string `json:"birthcity"  csv:"birthCity"  db:"birthCity"`
   Deathyear   string `json:"deathyear"  csv:"deathYear"  db:"deathYear"`
   Deathmonth   string `json:"deathmonth"  csv:"deathMonth"  db:"deathMonth"`
   Deathday   string `json:"deathday"  csv:"deathDay"  db:"deathDay"`
   Deathcountry   string `json:"deathcountry"  csv:"deathCountry"  db:"deathCountry"`
   Deathstate   string `json:"deathstate"  csv:"deathState"  db:"deathState"`
   Deathcity   string `json:"deathcity"  csv:"deathCity"  db:"deathCity"`
   Namefirst   string `json:"namefirst"  csv:"nameFirst"  db:"nameFirst"`
   Namelast   string `json:"namelast"  csv:"nameLast"  db:"nameLast"`
   Namegiven   string `json:"namegiven"  csv:"nameGiven"  db:"nameGiven"`
   Weight   string `json:"weight"  csv:"weight"  db:"weight"`
   Height   string `json:"height"  csv:"height"  db:"height"`
   Bats   string `json:"bats"  csv:"bats"  db:"bats"`
   Throws   string `json:"throws"  csv:"throws"  db:"throws"`
   Debut   string `json:"debut"  csv:"debut"  db:"debut"`
   Finalgame   string `json:"finalgame"  csv:"finalGame"  db:"finalGame"`
   Retroid   string `json:"retroid"  csv:"retroID"  db:"retroID"`
   Bbrefid   string `json:"bbrefid"  csv:"bbrefID"  db:"bbrefID"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *People) GetTableName() string {
  return "people"
}

// GetFileName returns the name of the source file the model was created from
func (m *People) GetFileName() string {
  return "People.csv"
}

// GetFilePath returns the path of the source file
func (m *People) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/People.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *People) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*People, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("People ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("People ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("People ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
