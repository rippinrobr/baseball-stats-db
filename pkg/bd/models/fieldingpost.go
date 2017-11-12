package models


type FieldingPost struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Pos   string `json:"pos"  csv:"POS"  db:"POS"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   string `json:"gs"  csv:"GS"  db:"GS"`
   Innouts   string `json:"innouts"  csv:"InnOuts"  db:"InnOuts"`
   Po   int `json:"po"  csv:"PO"  db:"PO"`
   A   int `json:"a"  csv:"A"  db:"A"`
   E   int `json:"e"  csv:"E"  db:"E"`
   Dp   int `json:"dp"  csv:"DP"  db:"DP"`
   Tp   int `json:"tp"  csv:"TP"  db:"TP"`
   Pb   string `json:"pb"  csv:"PB"  db:"PB"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
}

func (m *FieldingPost) GetTableName() string {
  return "FieldingPost"
}

func (m *FieldingPost) GetFileName() string {
  return "FieldingPost.csv"
}

func (m *FieldingPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/FieldingPost.csv"
}
