package liquor

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
	"github.com/aodin/denver/geocode"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// rawHearing is an unparsed hearing
type rawHearing struct {
	Link    string
	Name    string
	Address string
	Date    string
	Time    string
	Outcome string
}

var Hearings = aspect.Table("hearings",
	aspect.Column("id", aspect.Integer{PrimaryKey: true}),
	aspect.Column("notice_link", aspect.String{}),
	aspect.Column("name", aspect.String{}),
	aspect.Column("address", aspect.String{}),
	aspect.Column("latitude", aspect.Real{}),
	aspect.Column("longitude", aspect.Real{}),
	aspect.Column("time", aspect.Timestamp{}),
	aspect.Column("outcome", aspect.String{}),
	aspect.Column("location", postgis.Geometry{postgis.Point{}, 4326}),
)

// Hearing is a public hearing with a time and location
type Hearing struct {
	Id         int64     `json:"id"`
	NoticeLink string    `json:"notice_link"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Time       time.Time `json:"time"`
	Outcome    string    `json:"outcome"`
}

func ParseHearingsJSON(contents []byte) (hearings []Hearing, err error) {
	err = json.Unmarshal(contents, &hearings)
	return
}

var layout = `Jan 02, 2006 3:04 PM`
var urlRoot = `http://www.denvergov.org/HearingViewerApplication/`

func CleanHearings(raws []rawHearing, g geocode.Geocoder) ([]Hearing, error) {
	hearings := make([]Hearing, len(raws))
	for i, raw := range raws {
		hearing, err := raw.Convert(g)
		if err != nil {
			return hearings, fmt.Errorf("Error while convert hearing %d: %s", i, err)
		}
		hearings[i] = hearing
	}
	return hearings, nil
}

func (r rawHearing) Convert(g geocode.Geocoder) (h Hearing, err error) {
	// Copy the strings
	h.Name = r.Name
	h.Address = r.Address
	h.Outcome = r.Outcome

	// Create the complete link
	link := urlRoot + r.Link

	// Parse the link to get the id
	u, err := url.Parse(link)
	if err != nil {
		return
	}

	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}

	rawid := values.Get("id")
	h.Id, err = strconv.ParseInt(rawid, 10, 64)
	if err != nil {
		return
	}

	// Build the full notice link
	h.NoticeLink = u.String()

	// Get the Mountain timezone
	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		return
	}
	// Parse the time
	h.Time, err = time.ParseInLocation(layout, r.Date+" "+r.Time, loc)
	if err != nil {
		return
	}

	// Perform geolocation on the address
	h.Latitude, h.Longitude, err = g.Geocode(r.Address)
	return
}

func ParseHearingsHTML(content []byte) (hearings []rawHearing, err error) {
	// Parse the entire document
	d, err := gokogiri.ParseHtml(content)
	if err != nil {
		err = fmt.Errorf("Error parsing HTML: %s", err)
		return
	}

	// Get the table rows
	q := "//table[@id='GridViewLiquorNoticeHearingSchedule']//tr"
	rows, err := d.Search(q)
	if err != nil {
		return
	}

	if len(rows) < 2 {
		return
	}

	// Skip the first row in the table body - it is a header
	var cells []xml.Node
	for i, row := range rows[1:] {
		// Select the cells from the row
		cells, err = row.Search("./td")
		if err != nil {
			err = fmt.Errorf("Error finding cells in row %d: %s", i, err)
			return
		}

		if len(cells) < 5 {
			err = fmt.Errorf("Row %d does not have 5 cells", i)
			return
		}

		// Get the link
		link := cells[0].FirstChild().Attributes()["href"]
		h := rawHearing{
			Link:    link.String(),
			Name:    strings.TrimSpace(cells[0].Content()),
			Address: strings.TrimSpace(cells[1].Content()),
			Date:    strings.TrimSpace(cells[2].Content()),
			Time:    strings.TrimSpace(cells[3].Content()),
			Outcome: strings.TrimSpace(cells[4].Content()),
		}
		hearings = append(hearings, h)
	}
	return
}
