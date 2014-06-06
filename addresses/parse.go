package addresses

import (
	"github.com/aodin/csv2"
	"io"
)

func Parse(f io.ReadCloser) (addresses []Address, err error) {
	defer f.Close()

	// Create a new reader instance
	csvf := csv.NewReader(f)

	// Discard the header
	if _, err = csvf.Read(); err != nil {
		return
	}

	// Read the addresses
	err = csvf.Unmarshal(&addresses)
	return
}
