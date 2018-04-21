package retrosheet

import (
	"strconv"
	"time"
)

// The only 'data scrubbing' I'm doing is on the GameNumber property since the gamelog
// adds A and B to the GameNumber definition.  A is game one of a three team double header
// and will be changed to be 31 since it represents 3 team double header game 1. B will be
// converted to 32 since its the second game of a 3 game double header.
type gameUniqueID struct {
	Season       int        `json:"season" bson:"season" db:"season"`
	GameDateStr  string     `json:"gameDateStr" bson:"gameDateStr" db:"game_date_str"`
	GameMonth    time.Month `json:"gameMonth" bson:"gameMonth" db:"game_month"`
	GameMonthDay int        `json:"gameMonthDay" bson:"gameMonthDay" db:"game_month_day"`
	GameNumber   int        `json:"gameNumber" bson:"gameNumber" db:"game_number"`
}

// ConvertNonIntGameNumberValues checks to see if strGameNum is a string that can be
// converted to an int.  If the value of strGameNum is A then 31 will be returned. If
// the value is B then 32 is returned.
func convertNonIntGameNumberValues(strGameNum string) (int, error) {
	if strGameNum == "A" {
		return 31, nil
	}

	if strGameNum == "B" {
		return 32, nil
	}

	return strconv.Atoi(strGameNum)
}
