package models


type AwardsShareManagers struct {
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Pointswon   int `json:"pointswon"  csv:"pointsWon"  db:"pointsWon"`
   Pointsmax   int `json:"pointsmax"  csv:"pointsMax"  db:"pointsMax"`
   Votesfirst   int `json:"votesfirst"  csv:"votesFirst"  db:"votesFirst"`
}

func (m *AwardsShareManagers) GetTableName() string {
  return "AwardsShareManagers"
}

func (m *AwardsShareManagers) GetFileName() string {
  return "AwardsShareManagers.csv"
}

func (m *AwardsShareManagers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsShareManagers.csv"
}
