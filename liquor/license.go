package liquor

import (
	"fmt"
	"strconv"
	"time"
)

// License
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
// -------------

// Bounding box
// <westbc>-105.109336</westbc>
// <eastbc>-104.671208</eastbc>
// <northbc>39.863214</northbc>
// <southbc>39.614990</southbc>

// Categories, as of 2014-02-01
// * LI32 - LIQUOR-3.2 % BEER			178
// * LIRE - LIQUOR-RETAIL				204
// * LIBW - LIQUOR-BEER & WINE			 88
// * TAST - LIQUOR TASTING				 33
// * LITA - LIQUOR-TAVERN				254
// * LICL - LIQUOR-CLUB					 25
// * LIAR - LIQUOR-ARTS					 10
// * CAB8 - UNDERAGE CABARET PATRON		 21
// * LIHR - LIQUOR-HOTEL/RESTAURANT		817
// * CABA - CABARET						401
// * LIDR - LIQUOR-DRUG STORE			  3
// * LIBR - LIQUOR-BREW-PUB				  8

type License struct {
	UniqueId    string    `json:"id"`
	LicenseId   string    `json:"-"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Category    string    `json:"type"`
	LicenseName string    `json:"-"`
	Issued      time.Time `json:"issued"`
	Expires     time.Time `json:"expires"`
	Xcoord      float64   `json:"x"`
	Ycoord      float64   `json:"y"`
}

func (l *License) String() string {
	return fmt.Sprintf("%s, %s", l.Name, l.Address)
}

// Convert the record into a string array that can be written to a CSV
func (l *License) CSV() []string {
	return []string{
		l.UniqueId,
		l.LicenseId,
		l.Name,
		l.Address,
		l.Category,
		l.LicenseName,
		l.Issued.Format("2006-01-02"),
		l.Expires.Format("2006-01-02"),
		strconv.FormatFloat(l.Xcoord, 'f', 1, 64),
		strconv.FormatFloat(l.Ycoord, 'f', 1, 64),
	}
}

// Are the two licenses equal?
// TODO There's not a better way to do this - use reflect
func (l *License) Equals(other *License) bool {
	if l.UniqueId != other.UniqueId {
		return false
	}
	if l.LicenseId != other.LicenseId {
		return false
	}
	if l.Name != other.Name {
		return false
	}
	if l.Address != other.Address {
		return false
	}
	if l.Category != other.Category {
		return false
	}
	if l.LicenseName != other.LicenseName {
		return false
	}
	if l.Issued != other.Issued {
		return false
	}
	if l.Expires != other.Expires {
		return false
	}
	if l.Xcoord != other.Xcoord {
		return false
	}
	if l.Ycoord != other.Ycoord {
		return false
	}
	return true
}

type Change struct {
	prev interface{}
	cur  interface{}
}

// Save the changes between the licenses to a map
func (l *License) Changes(other *License) map[string]interface{} {
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
	if l.Xcoord != other.Xcoord {
		diff["Xcoord"] = Change{l.Xcoord, other.Xcoord}
	}
	if l.Ycoord != other.Ycoord {
		diff["Ycoord"] = Change{l.Ycoord, other.Ycoord}
	}
	return diff
}
