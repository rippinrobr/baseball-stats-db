package parsers

import (
	"errors"
	"os"

	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"
)

// ParseAndLoadCSV is responsible for parsing a CSV file and loading the array
// with structs that represent each line of the file.  The array parameter must
// be an array of pointers
func ParseAndLoadCSV(f *os.File, objs []models.TableObject) (models.TableObject, error) {
	return nil, errors.New("Not Implemented")
}
