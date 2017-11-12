package models


type CollegePlaying struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Schoolid   string `json:"schoolid"  csv:"schoolID"  db:"schoolID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
}

func (m *CollegePlaying) GetTableName() string {
  return "CollegePlaying"
}

func (m *CollegePlaying) GetFileName() string {
  return "CollegePlaying.csv"
}

func (m *CollegePlaying) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/CollegePlaying.csv"
}
