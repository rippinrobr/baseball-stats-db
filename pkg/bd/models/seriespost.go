package models


type SeriesPost struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Teamidwinner   string `json:"teamidwinner"  csv:"teamIDwinner"  db:"teamIDwinner"`
   Lgidwinner   string `json:"lgidwinner"  csv:"lgIDwinner"  db:"lgIDwinner"`
   Teamidloser   string `json:"teamidloser"  csv:"teamIDloser"  db:"teamIDloser"`
   Lgidloser   string `json:"lgidloser"  csv:"lgIDloser"  db:"lgIDloser"`
   Wins   int `json:"wins"  csv:"wins"  db:"wins"`
   Losses   int `json:"losses"  csv:"losses"  db:"losses"`
   Ties   int `json:"ties"  csv:"ties"  db:"ties"`
}

func (m *SeriesPost) GetTableName() string {
  return "SeriesPost"
}

func (m *SeriesPost) GetFileName() string {
  return "SeriesPost.csv"
}

func (m *SeriesPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/SeriesPost.csv"
}
