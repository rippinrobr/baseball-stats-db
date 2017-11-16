package csv

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// ParserFunc is a way to allow the parsing function  to be
// passed around to the Model's ParseCSV func so that I can
// pass it a 'real' slice of the TableObject Type
type ParserFunc func(*os.File, interface{}) error

// ParseAndLoad is responsible for parsing a CSV file and loading the array
// with structs that represent each line of the file.  The array parameter must
// be an array of pointers
func ParseAndLoad(f *os.File, r interface{}) (interface{}, error) {
	// fmt.Println(reflect.TypeOf(r))
	// rows := reflect.New(reflect.SliceOf(reflect.TypeOf(r))).Interface()
	// fmt.Println(rows, reflect.TypeOf(rows))
	// an example of how this can work
	// https://play.golang.org/p/-2DdzEgrhM

	if err := gocsv.UnmarshalFile(f, r); err != nil { // Load clients from file
		fmt.Println("parser error")
		return nil, err
	}

	for _, row := range r.([]interface{}) {
		fmt.Println(row)
	}

	return r, nil
	//return nil, errors.New("Not Implemented")
}
