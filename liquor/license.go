package liquor

import (
	"bytes"
	"fmt"
	"os/exec"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// License layout:
//  1. UNIQUE_ID
//  2. BFN
//  3. LIC_ID
//  4. BUS_PROF_NAME
//  5. FULL_ADDRESS
//  6. CODE						LI32 / LIRE / LIHR
//  7. CATEGORY
//  8. LIC_NAME					LIQUOR-RETAIL / LIQUOR-HOTEL/RESTAURANT
//  9. DESCRIPTION				beer & wine, hotel/restaurant
// 10. IDATE					2013-09-07 00:00:00
// 11. EDATE
// 12. LIC_STATUS
// 13. ADD_ID
// 14. EXTERNAL_ADDRESS_ID
// 15. POLICE_DIST
// 16. COUNCIL_DIST
// 17. CENSUS_TRACT
// 18. OVERRIDE
// 19. X_COORD					In CSV as floats, e.g.: 3115317.0
// 20. Y_COORD

type License struct {
	UniqueId    string    `db:"id" json:"id"`
	BFN         string    `db:"bfn" json:"-"`
	LicenseId   string    `db:"license" json:"-"`
	Name        string    `db:"name" json:"name"`
	Address     string    `db:"address" json:"address"`
	Code        string    `db:"code" json:"-"`
	Category    string    `db:"category" json:"type"`
	LicenseName string    `db:"license_name" json:"-"`
	Description string    `db:"description" json:"-"`
	Issued      time.Time `db:"issued" json:"issued" csv:"2006-01-02"`
	Expires     time.Time `db:"expires" json:"expires" csv:"2006-01-02"`
	Status      string    `db:"status" json:"status"`
	AddId       string    `db:"add_id" json:"-"`
	ExtAddId    string    `db:"external_address_id" json:"-"`
	Police      string    `db:"police_district" json:"-"`
	Council     string    `db:"council_district" json:"-"`
	Census      string    `db:"census_tract" json:"-"`
	Override    string    `db:"override" json:"-"`
	Xcoord      float64   `db:"longitude" json:"longitude"`
	Ycoord      float64   `db:"latitude" json:"latitude"`
}

func (l License) String() string {
	return fmt.Sprintf("%s, %s", l.Name, l.Address)
}

// Convert the record into a string array that can be written to a CSV
func (l License) CSV() []string {
	return []string{
		l.UniqueId,
		l.BFN,
		l.LicenseId,
		l.Name,
		l.Address,
		l.Code,
		l.Category,
		l.LicenseName,
		l.Description,
		l.Issued.Format("2006-01-02"),
		l.Expires.Format("2006-01-02"),
		l.Status,
		l.AddId,
		l.ExtAddId,
		l.Police,
		l.Council,
		l.Census,
		l.Override,
		strconv.FormatFloat(l.Xcoord, 'f', 1, 64),
		strconv.FormatFloat(l.Ycoord, 'f', 1, 64),
	}
}

// The normalized CSV output increases the number of significant digits for
// the x and y coords, since they are now latitude and longitudes
func (l License) NormalizedCSV() []string {
	return []string{
		l.UniqueId,
		l.BFN,
		l.LicenseId,
		l.Name,
		l.Address,
		l.Code,
		l.Category,
		l.LicenseName,
		l.Description,
		l.Issued.Format("2006-01-02"),
		l.Expires.Format("2006-01-02"),
		l.Status,
		l.AddId,
		l.ExtAddId,
		l.Police,
		l.Council,
		l.Census,
		l.Override,
		strconv.FormatFloat(l.Xcoord, 'f', 8, 64),
		strconv.FormatFloat(l.Ycoord, 'f', 8, 64),
	}
}

func NormalizedHeader() []string {
	return []string{
		"UniqueId",
		"BFN",
		"LicenseId",
		"Name",
		"Address",
		"Category",
		"Code",
		"LicenseName",
		"Description",
		"Issued",
		"Expires",
		"Status",
		"AddId",
		"ExtAddId",
		"Police",
		"Council",
		"Census",
		"Override",
		"Longitude",
		"Latitude",
	}
}

// Are the two licenses equal?
// TODO There's not a better way to do this - use reflect deep equals
func (l License) Equals(other License) bool {
	return reflect.DeepEqual(l, other)
}

type Change struct {
	prev interface{}
	cur  interface{}
}

// Save the changes between the licenses to a map
func (l License) Changes(other License) map[string]interface{} {
	diff := make(map[string]interface{})
	// The unique id should never change since we're using that for mapping
	if l.Name != other.Name {
		diff["Name"] = Change{l.Name, other.Name}
	}
	if l.Address != other.Address {
		diff["Address"] = Change{l.Address, other.Address}
	}
	if l.Category != other.Category {
		diff["Category"] = Change{l.Category, other.Category}
	}
	if l.Issued != other.Issued {
		diff["Issued"] = Change{l.Issued, other.Issued}
	}
	if l.Expires != other.Expires {
		diff["Expires"] = Change{l.Expires, other.Expires}
	}
	if l.Status != other.Status {
		diff["Status"] = Change{l.Status, other.Status}
	}
	if l.Xcoord != other.Xcoord {
		diff["Xcoord"] = Change{l.Xcoord, other.Xcoord}
	}
	if l.Ycoord != other.Ycoord {
		diff["Ycoord"] = Change{l.Ycoord, other.Ycoord}
	}
	return diff
}

// Convert an array of licenses to latitude and longitude
func StatePlaneToLatLong(licenses []License) ([]License, error) {
	converted := make([]License, len(licenses))

	coords := make([]string, len(licenses))
	for i, license := range licenses {
		coords[i] = fmt.Sprintf("%f %f", license.Xcoord, license.Ycoord)
	}

	cmd := exec.Command(
		"gdaltransform",
		"-s_srs",
		"EPSG:2232",
		"-t_srs",
		"EPSG:4326",
	)

	// Create a single argument string
	cmd.Stdin = strings.NewReader(strings.Join(coords, "\n"))

	// Return as a bytes buffer
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return converted, err
	}

	// The command returns as an x, y, and z coordinate
	xyzs := strings.Split(out.String(), "\n")

	var lnglat []string
	var lat, lng float64
	for i, license := range licenses {
		// Remember that longitude is an x coordinate!
		lnglat = strings.SplitN(xyzs[i], " ", 3)

		lat, err = strconv.ParseFloat(lnglat[1], 64)
		if err != nil {
			return converted, err
		}
		lng, err = strconv.ParseFloat(lnglat[0], 64)
		if err != nil {
			return converted, err
		}

		license.Xcoord = lng
		license.Ycoord = lat
		converted[i] = license
	}
	return converted, nil
}

func ById(ls []License) (byId map[string]License) {
	byId = make(map[string]License)
	for _, license := range ls {
		_, exists := byId[license.UniqueId]
		if exists {
			// log.Printf("Unique Id %s already exists on line %d\n", license.UniqueId, i + 2)
			// TODO Are the licenses the same?
		} else {
			byId[license.UniqueId] = license
		}
	}
	return
}

// Normalize will:
// * Remove duplicate unique ids
// * Convert the Colorado state plane coordinates to latitude and longitude
// * Sort the licenses by unique id
func Normalize(originals []License) ([]License, error) {
	// Convert to a map to remove duplicates
	// TODO This should error if duplicates aren't equal
	mapping := ById(originals)

	licenses := make([]License, len(mapping))
	var index int
	for _, license := range mapping {
		licenses[index] = license
		index += 1
	}

	// Convert the x and y coords to latitude and longitude
	converted, err := StatePlaneToLatLong(licenses)
	if err != nil {
		return nil, err
	}

	// Sort by unique id
	OrderedLicenses(converted).Sort()

	return converted, nil
}

type OrderedLicenses []License

// Implement the sort.Interface for sorting
func (a OrderedLicenses) Len() int {
	return len(a)
}

func (a OrderedLicenses) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByUniqueId struct {
	OrderedLicenses
}

// Sort by the unique Id string
func (a ByUniqueId) Less(i, j int) bool {
	return a.OrderedLicenses[i].UniqueId > a.OrderedLicenses[j].UniqueId
}

func (a OrderedLicenses) Sort() {
	sort.Sort(ByUniqueId{a})
}
