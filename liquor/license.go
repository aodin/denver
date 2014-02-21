package liquor

import (
	"fmt"
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
	Latitude    float64   `json:"lat"`
	Longitude   float64   `json:"long"`
}

func (l *License) String() string {
	return fmt.Sprintf("%s, %s", l.Name, l.Address)
}

// Are the two licenses equal?
func (l *License) Equals(other *License) bool {
	return false
}
