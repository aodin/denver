package geocode

import (
	"testing"
)

func TestGoogle(t *testing.T) {
	g := "1062 Delaware St"
	lat, lng, err := Google.Geocode(g)
	if err != nil {
		t.Fatal(err)
	}

	// TODO Lat and long accuracy should be a lossy check
	if lat != 39.733536 {
		t.Errorf("Unexpected lat: %f", lat)
	}
	if lng != -104.992611 {
		t.Errorf("Unexpected lng: %f", lng)
	}

	// The query should not have modified the internal Google URL
	x := `https://maps.googleapis.com/maps/api/geocode/json`
	if Google.url.String() != x {
		t.Errorf("Google URL was modified: %s", Google.url.String())
	}
}
