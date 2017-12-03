package csv

import (
	"os"
)

// ParserFunc is a way to allow the parsing function  to be
// passed around to the Model's ParseCSV func so that I can
// pass it a 'real' slice of the TableObject Type
type ParserFunc func(*os.File, interface{}) error
