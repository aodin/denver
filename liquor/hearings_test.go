package liquor

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type testGeocoder struct{}

func (g testGeocoder) Geocode(address string) (lat, lng float64, err error) {
	lat, lng = 1.0, 1.0
	return
}

var example = rawHearing{
	Name:    "7 - Eleven Store 37016a",
	Address: "4922 N Willow St",
	Date:    "Jun 25, 2014",
	Time:    "9:00 AM",
	Outcome: "Pending",
}

func TestConvert(t *testing.T) {
	hearing, err := example.Convert(testGeocoder{})
	if err != nil {
		t.Fatalf("Error during raw hearing conversion: %s", err)
	}

	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		t.Fatalf("could not load Mountain timezone")
	}
	x := time.Date(2014, time.June, 25, 9, 0, 0, 0, loc).String()
	// TODO Why do I have to call Stirng()?
	if hearing.Time.String() != x {
		t.Errorf("Unexpected time: %s != %s", hearing.Time, x)
	}

	if hearing.Name != "7 - Eleven Store 37016a" {
		t.Errorf("Unexpected name: %s", hearing.Name)
	}
	if hearing.Address != "4922 N Willow St" {
		t.Errorf("Unexpected address: %s", hearing.Address)
	}
	if hearing.Outcome != "Pending" {
		t.Errorf("Unexpected outcome: %s", hearing.Name)
	}
	if hearing.Latitude != 1.0 {
		t.Errorf("Unexpected lat: %f", hearing.Latitude)
	}
	if hearing.Longitude != 1.0 {
		t.Errorf("Unexpected lng: %f", hearing.Longitude)
	}
}

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

	// Examine the first hearing
	h := hearings[0]
	if h.Name != "7 - Eleven Store 37016a" {
		t.Errorf("Unexpected parsed name: %s", h.Name)
	}
	if h.Address != "4922 N Willow St" {
		t.Errorf("Unexpected parsed address: %s", h.Address)
	}
	if h.Date != "Jun 25, 2014" {
		t.Errorf("Unexpected parsed date: %s", h.Date)
	}
	if h.Time != "9:00 AM" {
		t.Errorf("Unexpected parsed time: %s", h.Time)
	}
	if h.Outcome != "Pending" {
		t.Errorf("Unexpected parsed outcome: %s", h.Name)
	}
}
