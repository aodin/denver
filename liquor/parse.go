package liquor

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	UnexpectedLength     = errors.New("liquor: unexpected csv row length")
	UnparsableTimestamp  = errors.New("liquor: unparsable timestamp")
	UnparsableCoordinate = errors.New("liquor: unparsable coordinate")
)

// TODO Assume that timestamps are mountain time?
func ParseTimestamp(input string) (time.Time, error) {
	// In the absence of a time zone indicator, Parse returns a time in UTC.
	return time.Parse("2006-01-02 15:04:05", input)
}

func ParseLicense(raw []string) (License, error) {
	if len(raw) != 20 {
		return License{}, UnexpectedLength
	}
	l := License{
		UniqueId:    raw[0],
		BFN:         raw[1],
		LicenseId:   raw[2],
		Name:        strings.TrimSpace(raw[3]),
		Address:     strings.TrimSpace(raw[4]),
		Category:    strings.TrimSpace(raw[5]),
		Code:        strings.TrimSpace(raw[6]),
		LicenseName: strings.TrimSpace(raw[7]),
		Description: strings.TrimSpace(raw[8]),
		Status:      strings.TrimSpace(raw[11]),
		AddId:       strings.TrimSpace(raw[12]),
		ExtAddId:    strings.TrimSpace(raw[13]),
		Police:      strings.TrimSpace(raw[14]),
		Council:     strings.TrimSpace(raw[15]),
		Census:      strings.TrimSpace(raw[16]),
		Override:    strings.TrimSpace(raw[17]),
	}

	var err error
	// Parse the Issued and Expires timestamps
	l.Issued, err = ParseTimestamp(raw[9])
	if err != nil {
		return l, UnparsableTimestamp
	}

	l.Expires, err = ParseTimestamp(raw[10])
	if err != nil {
		return l, UnparsableTimestamp
	}

	// Colorado state plane coordinates for now
	l.Xcoord, err = strconv.ParseFloat(raw[18], 64)
	if err != nil {
		return l, UnparsableCoordinate
	}
	l.Ycoord, err = strconv.ParseFloat(raw[19], 64)
	if err != nil {
		return l, UnparsableCoordinate
	}
	return l, nil
}

func ParseLicensesCSV(path string) ([]License, error) {
	licenses := make([]License, 0)
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
