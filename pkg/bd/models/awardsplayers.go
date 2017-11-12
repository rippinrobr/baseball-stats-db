package models


type AwardsPlayers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Tie   string `json:"tie"  csv:"tie"  db:"tie"`
   Notes   string `json:"notes"  csv:"notes"  db:"notes"`
}

func (m *AwardsPlayers) GetTableName() string {
  return "AwardsPlayers"
}

func (m *AwardsPlayers) GetFileName() string {
  return "AwardsPlayers.csv"
}

func (m *AwardsPlayers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsPlayers.csv"
}
