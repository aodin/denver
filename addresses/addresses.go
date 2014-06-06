package addresses

import ()

type Address struct {
	Type              string
	X                 float64 // Likely the Colorado State Plane
	Y                 float64 // Likely the Colorado State Plane
	Latitude          float64
	Longitude         float64
	NumberPrefix      string
	Number            string
	NumberSuffix      string
	PreModifier       string
	PreDirection      string
	Street            string
	StreetType        string
	PostDirection     string
	PostModifier      string
	BuildingType      string
	BuildingId        string
	UnitType          string
	UnitNumber        string
	CompositeUnitType string
	CompositeUnitId   string
	Full              string
}

func (a Address) String() string {
	return a.Full
}

func (a Address) LatLong() (float64, float64) {
	return a.Latitude, a.Longitude
}
