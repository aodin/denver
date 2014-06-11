package liquor

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseHTML(t *testing.T) {
	f, err := os.Open("./hearings_2014-06-11.html")
	if err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	hearings, err := ParseHTML(content)
	if err != nil {
		t.Fatal(err)
	}

	if len(hearings) != 21 {
		t.Fatalf("Unexpected length of hearings: %d", len(hearings))
	}

	h := hearings[0]

	if h.Name != "7 - Eleven Store 37016a" {
		t.Fatalf("Unexpected name of hearing: %s", h.Name)
	}
	if h.Address != "4922 N Willow St" {
		t.Fatalf("Unexpected address of hearing: %s", h.Address)
	}
	if h.Date != "Jun 25, 2014" {
		t.Fatalf("Unexpected date of hearing: %s", h.Date)
	}
	if h.Time != "9:00 AM" {
		t.Fatalf("Unexpected time of hearing: %s", h.Time)
	}
	if h.Outcome != "Pending" {
		t.Fatalf("Unexpected outcome of hearing: %s", h.Name)
	}
}
