package crime

import (
	"testing"
)

func TestParseOffenseCodesCSV(t *testing.T) {
	raw, err := ParseOffenseCodesCSV("example_codes.csv")
	if err != nil {
		t.Fatalf("Error during ParseOffenseCodesCSV: %s", err)
	}
	if len(raw) != 9 {
		t.Fatalf("Unexpected length of raw codes: %d", len(raw))
	}

	// Convert to codes
	codes := ConvertRawCodes(raw)
	stolen := codes[0]

	output := stolen.String()
	expected := `2804-1: Possession of stolen property (All Other Crimes)`

	if output != expected {
		t.Fatalf("Unexpected string output of a code: %s", output)
	}
}

func TestParseCrimeCSV(t *testing.T) {
	raw, err := ParseCrimeCSV("example_crimes.csv")
	if err != nil {
		t.Fatalf("Error during ParseCrimeCSV: %s", err)
	}
	if len(raw) != 7 {
		t.Fatalf("Unexpected length of raw crimes: %d", len(raw))
	}

	// Convert to crimes
	crimes, err := ConvertRawCrimes(raw)
	if err != nil {
		t.Fatalf("Unexpected error during convert crimes: %d", len(crimes))
	}
	if len(crimes) != 7 {
		t.Fatalf("Unexpected length of converted crimes: %d", len(crimes))
	}
}
