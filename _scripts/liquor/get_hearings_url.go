package main

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/denver/geocode"
	"github.com/aodin/denver/liquor"
	"io/ioutil"
	"net/http"
)

// Get the latest liquor hearings from the inner frame of:
// http://www.denvergov.org/businesslicensing/DenverBusinessLicensingCenter/PublicHearingSchedule/tabid/441585/Default.aspx

// Which is
// http://www.denvergov.org/HearingViewerApplication/default.aspx

var url = `http://www.denvergov.org/HearingViewerApplication/default.aspx`

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Parse the contents
	raws, err := liquor.ParseHearingsHTML(contents)
	if err != nil {
		panic(err)
	}

	hearings, err := liquor.CleanHearings(raws, geocode.Google)
	if err != nil {
		panic(err)
	}

	// Pretty print the output!
	b, err := json.MarshalIndent(hearings, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)
}
