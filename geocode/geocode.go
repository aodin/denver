package geocode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Geocoder interface {
	Geocode(address string) (lat, lng float64, err error)
}

// Users of the free API:
// * 2,500 requests per 24 hour period.
// * 10 requests per second.
type google struct {
	rate    int64
	rawurl  string
	limiter <-chan time.Time
	url     url.URL
}

func (g google) getURL() url.URL {
	return g.url
}

type response struct {
	Results []result `json:"results"`
	Status  string   `json:"status"`
}

type result struct {
	Address  string   `json:"formatted_address"`
	Geometry geometry `json:"geometry"`
}

type geometry struct {
	Location location `json:"location"`
}

type location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geocode returns the latitude and longitude of an address. "Denver
// Colorado" is automatically added to the query.
func (g *google) Geocode(address string) (lat, lng float64, err error) {
	// Wait for a tick
	<-g.limiter

	// Adds Denver
	full := address + " Denver Colorado"

	// TODO components?
	// https://developers.google.com/maps/documentation/geocoding/#ComponentFiltering
	v := url.Values{}
	v.Set("address", full)
	u := g.getURL()
	u.RawQuery = v.Encode()

	// Get the JSON from google
	r, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer r.Body.Close()
	contents, err := ioutil.ReadAll(r.Body)
	var resp response
	if err = json.Unmarshal(contents, &resp); err != nil {
		return
	}
	if resp.Status != "OK" {
		err = fmt.Errorf("Unexpected status: %s", resp.Status)
		return
	}

	// There should be at least one result
	if len(resp.Results) < 1 {
		err = fmt.Errorf("Zero results returned")
		return
	}

	loc := resp.Results[0].Geometry.Location
	lat = loc.Lat
	lng = loc.Lng
	return
}

var Google = &google{
	rate:   9,
	rawurl: `https://maps.googleapis.com/maps/api/geocode/json`,
}

func init() {
	// Start the ticker
	Google.limiter = time.Tick(time.Duration(int64(time.Second) / Google.rate))

	// Parse the url
	u, err := url.Parse(Google.rawurl)
	if err != nil {
		panic(err)
	}

	Google.url = *u
}
