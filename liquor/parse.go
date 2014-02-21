package liquor

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

var (
	UnexpectedLength     = errors.New("liquor: unexpected csv row length")
	UnparsableTimestamp  = errors.New("liquor: unparsable timestamp")
	UnparsableCoordinate = errors.New("liquor: unparsable coordinate")
)

func ParseTimestamp(input string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", input)
}

func ParseLicense(raw []string) (*License, error) {
	if len(raw) != 20 {
		return nil, UnexpectedLength
	}
	license := &License{
		UniqueId:    raw[0],
		LicenseId:   raw[2],
		Name:        strings.TrimSpace(raw[3]),
		Address:     strings.TrimSpace(raw[4]),
		Category:    strings.TrimSpace(raw[5]),
		LicenseName: strings.TrimSpace(raw[7]),
	}

	var err error
	// Parse the Issued and Expires timestamps
	license.Issued, err = ParseTimestamp(raw[9])
	if err != nil {
		return nil, UnparsableTimestamp
	}

	license.Expires, err = ParseTimestamp(raw[10])
	if err != nil {
		return nil, UnparsableTimestamp
	}

	// TODO Ignore the colorado state plane coordinates for now
	return license, nil
}

func ParseLicensesCSV(path string) ([]*License, error) {
	licenses := make([]*License, 0)
	file, err := os.Open(path)
	if err != nil {
		return licenses, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	// Skip the header
	_, err = r.Read()
	if err != nil {
		return licenses, err
	}
	// TODO Re-examine the csv.ReadAll() source
	for {
		line, err := r.Read()
		if err == io.EOF {
			// EOF is an expected error!
			return licenses, nil
		}
		if err != nil {
			return licenses, err
		}
		license, err := ParseLicense(line)
		if err != nil {
			return licenses, err
		}
		licenses = append(licenses, license)
	}

}
