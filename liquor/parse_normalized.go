package liquor

import (
	"github.com/aodin/csv2"
	"os"
)

func ParseNormalizedLicensesCSV(path string) (licenses []License, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	r := csv.NewReader(file)

	// Skip the header
	_, err = r.Read()
	if err != nil {
		return
	}

	err = r.Unmarshal(&licenses)
	return
}
