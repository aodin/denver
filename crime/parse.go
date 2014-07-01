package crime

import (
	"fmt"
	"github.com/aodin/csv2"
	"os"
	"time"
)

var layout = "2006-01-02 15:04:05"

func ParseCrimeCSV(path string) (crimes []rawCrime, err error) {
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

	err = r.Unmarshal(&crimes)
	return
}

func ConvertRawCrime(raw rawCrime, l *time.Location) (crime Crime, err error) {
	// Composite field
	crime.CodeID = fmt.Sprintf("%d-%d", raw.Code, raw.CodeExt)

	// Fields that can be copied
	crime.IncidentID = raw.IncidentID
	crime.OffenseID = raw.OffenseID
	crime.Code = raw.Code
	crime.CodeExt = raw.CodeExt
	crime.Type = raw.Type
	crime.Category = raw.Category

	// Parse the times in the Denver timezone
	if crime.FirstOccurrence, err = time.ParseInLocation(layout, raw.FirstOccurrence, l); err != nil {
		return
	}
	var last time.Time
	if raw.LastOccurrence != "" {
		if last, err = time.ParseInLocation(layout, raw.LastOccurrence, l); err != nil {
			return
		}
		crime.LastOccurrence = &last
	}
	if crime.Reported, err = time.ParseInLocation(layout, raw.Reported, l); err != nil {
		return
	}

	crime.Address = raw.Address
	crime.Latitude = raw.Latitude
	crime.Longitude = raw.Longitude
	crime.District = raw.District
	crime.Precinct = raw.Precinct
	crime.Neighborhood = raw.Neighborhood
	return
}

func ConvertRawCrimes(raws []rawCrime) ([]Crime, error) {
	// Get the mountain timezone
	crimes := make([]Crime, len(raws))
	denver, tzErr := time.LoadLocation("America/Denver")
	if tzErr != nil {
		return crimes, tzErr
	}
	var err error
	for i, raw := range raws {
		if crimes[i], err = ConvertRawCrime(raw, denver); err != nil {
			return crimes, err
		}
	}
	return crimes, nil
}

func ParseOffenseCodesCSV(path string) (codes []rawCode, err error) {
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

	err = r.Unmarshal(&codes)
	return
}

func ConvertRawCode(raw rawCode) (code Code) {
	// Composite field
	code.ID = fmt.Sprintf("%d-%d", raw.Code, raw.Extension)

	// Fields that can be copied
	code.Code = raw.Code
	code.Extension = raw.Extension
	code.Description = raw.TypeName
	code.Category = raw.CategoryName
	if raw.IsCrime == 1 {
		code.IsCrime = true
	}
	if raw.IsTraffic == 1 {
		code.IsTraffic = true
	}
	return
}

func ConvertRawCodes(raws []rawCode) []Code {
	codes := make([]Code, len(raws))
	for i, raw := range raws {
		codes[i] = ConvertRawCode(raw)
	}
	return codes
}
