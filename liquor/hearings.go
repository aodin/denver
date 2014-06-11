package liquor

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"strings"
	"time"
)

type rawHearing struct {
	Name    string
	Address string
	Date    string
	Time    string
	Outcome string
}

// TODO Perform geolocation
type Hearing struct {
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
	Time      time.Time
	Outcome   string
}

func ParseHTML(content []byte) (hearings []rawHearing, err error) {
	// Parse the entire document
	d, err := gokogiri.ParseHtml(content)
	if err != nil {
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
	for _, row := range rows[1:] {
		// Select the cells from the row
		cells, err = row.Search("./td")
		if err != nil {
			return
		}

		if len(cells) < 5 {
			return
		}
		h := rawHearing{
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
