package models


type Salaries struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Salary   int `json:"salary"  csv:"salary"  db:"salary"`
}

func (m *Salaries) GetTableName() string {
  return "Salaries"
}

func (m *Salaries) GetFileName() string {
  return "Salaries.csv"
}

func (m *Salaries) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Salaries.csv"
}
