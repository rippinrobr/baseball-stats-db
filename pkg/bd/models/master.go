package models


type Master struct {
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

func (m *Master) GetTableName() string {
  return "Master"
}

func (m *Master) GetFileName() string {
  return "Master.csv"
}

func (m *Master) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Master.csv"
}
