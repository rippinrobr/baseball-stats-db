package models


type Fielding struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Pos   string `json:"pos"  csv:"POS"  db:"POS"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   string `json:"gs"  csv:"GS"  db:"GS"`
   Innouts   string `json:"innouts"  csv:"InnOuts"  db:"InnOuts"`
   Po   string `json:"po"  csv:"PO"  db:"PO"`
   A   string `json:"a"  csv:"A"  db:"A"`
   E   string `json:"e"  csv:"E"  db:"E"`
   Dp   string `json:"dp"  csv:"DP"  db:"DP"`
   Pb   string `json:"pb"  csv:"PB"  db:"PB"`
   Wp   string `json:"wp"  csv:"WP"  db:"WP"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Zr   string `json:"zr"  csv:"ZR"  db:"ZR"`
}

func (m *Fielding) GetTableName() string {
  return "Fielding"
}

func (m *Fielding) GetFileName() string {
  return "Fielding.csv"
}

func (m *Fielding) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Fielding.csv"
}
