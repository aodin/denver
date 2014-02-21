package addresses

import ()

type Address struct {
	Latitude   float64
	Longitude  float64
	Full       string
	Number     int
	PreDir     string
	Street     string
	PostType   string
	PostDir    string
	UnitType   string
	UnitNumber string
}

func (a Address) String() string {
	return a.Full
}

func (a Address) LatLong() (float64, float64) {
	return a.Latitude, a.Longitude
}
