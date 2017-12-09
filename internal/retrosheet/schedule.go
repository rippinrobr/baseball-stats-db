package retrosheet

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

//Schedule models a record in retrosheet.org *.SKED.TXT files with an additional Season int field
// GameType is one of these values Day (D), Night (N), Afternoon (A), Evening (E for twinight)
// PostponedOrCanceled - This field will contain one or more phrases related to the game if it was not played as scheduled. If there is more than one phrase, they are separated by a semi-colon (“;”). There are three possible outcomes for games not played on the originally scheduled date:
// ·         The game was played on another date
// ·         The game was played on the original date but at another site
// ·         The game was not played
// ·         This field can contain one or more standard phrases:
// ·         No Makeup played
// ·         Site Change
// ·         Played at XXX (where XXX is team/city code)
// ·         Rain
// ·         Cold
// ·         Snow
// ·         Threatening weather
// ·         Inclement weather
// ·         Hurricane XXX (where XXX is the name of the hurricane)
// ·         Wet grounds
// ·         Darkness
// ·         Power failure
// ·         Lighting failure
// ·         Schedule change
// ·         World War I
// ·         9/11 attack
// ·         Stadium … (damage or others words follow)
// ·         Strike
// ·         Sunday game stopped
// ·         Sunday curfew
// ·         Accommodate large crowd
// ·         XXX folded (where XXX is team code)
// ·         XXX played as YYY (where XXX & YYY are team codes)
// ·         XXX home team (where XXX is team code; used for games at alternate site)
// ·         Death of …
// ·         Funeral of …
// ·         Unknown reason
// MakeUpDate a  field will contain a makeup date if the postponed game was played at another time or place. If an attempt was known to have been made on a date but postponed again, that date will be listed. In that case, there will be a second date for the date when the game was actually played, in this form: “20150428; 20150528” For the note about a team folding, the team code is one of the standard Retrosheet team IDs. The same is true for the team played as note. Often, the two of these are combined in one field. The teams that folded are:
// 	Team/Lg          Last Game
// 	SR1 NL           9/10/1879
// 	TRN NL          9/20/1879
// 	ALT UA         5/31/1884 [moved to KCU 6/07]
// 	CHU UA        9/18/1884 [moved to MLU 9/27]
// 	PHU UA         8/07/1884 [moved to WIL 8/18 to 9/12; moved to SPU 9/27]
// 	WS7 AA         8/02/1884 [moved to RIC 8/05]
// 	BR4 AA          8/25/1890 [moved to BL3 8/27]
// 	CN3 AA         8/16/1891 [moved to ML3 8/18]
type Schedule struct {
	Season                   int          `json:"season" bson:"season" db:"season"`
	GameDateStr              string       `json:"gameDateStr" bson:"gameDateStr" db:"game_date_str"`
	GameMonth                time.Month   `json:"gameMonth" bson:"gameMonth" db:"game_month"`
	GameMonthDay             int          `json:"gameMonthDay" bson:"gameMonthDay" db:"game_month_day"`
	GameNumber               int          `json:"gameNumber" bson:"gameNumber" db:"game_number"`
	WeekDay                  time.Weekday `json:"weekDay" bson:"weekDay" db:"week_day"`
	VisitorsTeam             string       `json:"visitorsTeam" bson:"visitorsTeam" db:"visitors_team"`
	Visitorsleague           string       `json:"visitorsleague" bson:"visitorsleague" db:"visitors_league"`
	VisitorsSeasonGameNumber int          `json:"visitorsSeasonGameNumber" bson:"visitorsSeasonGameNumber" db:"visitors_season_game_number"`
	HomeTeam                 string       `json:"homeTeam" bson:"homeTeam" db:"home_team"`
	HomeLeague               string       `json:"homeLeague" bson:"homeLeague" db:"home_league"`
	HomeSeasonGameNumber     int          `json:"homeSeasonGameNumber" bson:"homeSeasonGameNumber" db:"home_season_game_number"`
	GameType                 string       `json:"gameType" bson:"gameType" db:"game_type"`
	PostponedOrCanceled      string       `json:"postponedOrCanceled" bson:"postponedOrCanceled" db:"postponed_canceled" `
	MakeUpDate               string       `json:"makeUpDate" bson:"makeUpDate" db:"make_up_date"`
}

var (
	monthMap = map[string]time.Month{
		"01": time.January,
		"02": time.February,
		"03": time.March,
		"04": time.April,
		"05": time.May,
		"06": time.June,
		"07": time.July,
		"08": time.August,
		"09": time.September,
		"10": time.October,
		"11": time.November,
		"12": time.December,
	}

	weekDayMap = map[string]time.Weekday{
		"Sun": time.Sunday,
		"Mon": time.Monday,
		"Tue": time.Tuesday,
		"Wed": time.Wednesday,
		"Thu": time.Thursday,
		"Fri": time.Friday,
		"Sat": time.Saturday,
	}
)

// NewSchedule create a new Schedule structure takes in the season and
// a slice of strings that represents a line in the *SKED.TXT file.
func NewSchedule(season int, data []string) (Schedule, error) {
	var err error
	s := Schedule{}

	if season < 1877 || season > time.Now().Year() {
		return s, errors.New("season must between than 1876 and the current year")
	}
	s.Season = season

	if len(data) != 12 {
		return s, fmt.Errorf("season data has 12 columns data only has %d", len(data))
	}

	s.GameDateStr = data[0]
	s.GameMonth = monthMap[s.GameDateStr[4:6]]
	s.GameMonthDay, _ = strconv.Atoi(s.GameDateStr[6:8])
	s.GameNumber, err = strconv.Atoi(data[1])
	if err != nil {
		log.Printf("unable to convert GameNumber value '%s'\n", data[1])
	}

	s.WeekDay = weekDayMap[data[2]]
	s.VisitorsTeam = data[3]
	s.Visitorsleague = data[4]
	s.VisitorsSeasonGameNumber, _ = strconv.Atoi(data[5])
	s.HomeTeam = data[6]
	s.HomeLeague = data[7]
	s.HomeSeasonGameNumber, _ = strconv.Atoi(data[8])
	s.GameType = data[9]
	s.PostponedOrCanceled = data[10]
	s.MakeUpDate = data[11]

	return s, nil
}
