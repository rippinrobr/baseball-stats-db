package models


type HomeGames struct {
   Yearkey   int `json:"yearkey"  csv:"year.key"  db:"year.key"`
   Leaguekey   string `json:"leaguekey"  csv:"league.key"  db:"league.key"`
   Teamkey   string `json:"teamkey"  csv:"team.key"  db:"team.key"`
   Parkkey   string `json:"parkkey"  csv:"park.key"  db:"park.key"`
   Spanfirst   string `json:"spanfirst"  csv:"span.first"  db:"span.first"`
   Spanlast   string `json:"spanlast"  csv:"span.last"  db:"span.last"`
   Games   int `json:"games"  csv:"games"  db:"games"`
   Openings   int `json:"openings"  csv:"openings"  db:"openings"`
   Attendance   int `json:"attendance"  csv:"attendance"  db:"attendance"`
}

func (m *HomeGames) GetTableName() string {
  return "HomeGames"
}

func (m *HomeGames) GetFileName() string {
  return "HomeGames.csv"
}

func (m *HomeGames) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/HomeGames.csv"
}
