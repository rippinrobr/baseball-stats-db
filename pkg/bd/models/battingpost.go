package models


type BattingPost struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ab   int `json:"ab"  csv:"AB"  db:"AB"`
   R   int `json:"r"  csv:"R"  db:"R"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Doubles   int `json:"doubles"  csv:"2B"  db:"2B"`
   Triples   int `json:"triples"  csv:"3B"  db:"3B"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Rbi   int `json:"rbi"  csv:"RBI"  db:"RBI"`
   Sb   int `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   int `json:"so"  csv:"SO"  db:"SO"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

func (m *BattingPost) GetTableName() string {
  return "BattingPost"
}

func (m *BattingPost) GetFileName() string {
  return "BattingPost.csv"
}

func (m *BattingPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/BattingPost.csv"
}
