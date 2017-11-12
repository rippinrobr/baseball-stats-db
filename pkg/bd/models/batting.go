package models


type Batting struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ab   string `json:"ab"  csv:"AB"  db:"AB"`
   R   string `json:"r"  csv:"R"  db:"R"`
   H   string `json:"h"  csv:"H"  db:"H"`
   Doubles   string `json:"doubles"  csv:"2B"  db:"2B"`
   Triples   string `json:"triples"  csv:"3B"  db:"3B"`
   Hr   string `json:"hr"  csv:"HR"  db:"HR"`
   Rbi   string `json:"rbi"  csv:"RBI"  db:"RBI"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Bb   string `json:"bb"  csv:"BB"  db:"BB"`
   So   string `json:"so"  csv:"SO"  db:"SO"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

func (m *Batting) GetTableName() string {
  return "Batting"
}

func (m *Batting) GetFileName() string {
  return "Batting.csv"
}

func (m *Batting) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Batting.csv"
}
