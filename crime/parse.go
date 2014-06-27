package crime

import (
	"fmt"
	"github.com/aodin/csv2"
	"os"
)

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
