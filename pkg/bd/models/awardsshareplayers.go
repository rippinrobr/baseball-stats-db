package models


type AwardsSharePlayers struct {
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Pointswon   int `json:"pointswon"  csv:"pointsWon"  db:"pointsWon"`
   Pointsmax   int `json:"pointsmax"  csv:"pointsMax"  db:"pointsMax"`
   Votesfirst   string `json:"votesfirst"  csv:"votesFirst"  db:"votesFirst"`
}

func (m *AwardsSharePlayers) GetTableName() string {
  return "AwardsSharePlayers"
}

func (m *AwardsSharePlayers) GetFileName() string {
  return "AwardsSharePlayers.csv"
}

func (m *AwardsSharePlayers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsSharePlayers.csv"
}
