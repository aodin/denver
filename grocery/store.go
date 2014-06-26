package grocery

import ()

// rawStore is a row parsed directly from the open data CSV file
type rawStore struct {
	Phone            string
	Hours            string
	Snap             string
	SIC              string
	NAICS            string
	Sales            string
	BranchStatus     string
	Employees        string
	EmployeeCategory string
	Franchise        string
	SquareFeet       string
	Source           string
	Type             string
	Longitude        string
	Latitude         string
	Name             string
	AddressID        string
	Address          string
	Address2         string
	City             string
	State            string
	ZIP              string
	GlobalID         string
}

// Various metadata information for grocery stores
var branch = map[string]string{
	"":  "unknown",
	"0": "unknown",
	"1": "Headquarters",
	"2": "Branch",
	"3": "Subsidiary Headquarters",
}

var sqrt = map[string]string{
	"":  "unknown",
	"A": "1 - 2499",
	"B": "2500 - 9999",
	"C": "10000 - 39999",
	"D": "40000+",
}

var types = map[string]string{
	"SM":   "Supermarket",
	"SGS":  "Small Grocery Store",
	"CS":   "Convenience store without gasoline",
	"CSG":  "Convenience store with gasoline",
	"SPCL": "Specialized foodstore",
	"WCS":  "Warehouse Club store",
	"SC":   "Supercenter/superstore",
	"DS":   "Dollar store",
	"PMCY": "Pharmacy/drugstore",
	"DMM":  "Discount mass-merchandise/department store",
}

type Store struct {
	Name          string
	Type          string
	Phone         string
	Hours         string
	AcceptsSnap   bool
	SIC           int64
	NAICS         int64
	SalesVolume   int64
	BranchStatus  string
	Employees     int64
	Franchise     string
	SquareFootage string
	Latitude      float64
	Longitude     float64
	Address       string
	AddressLine1  string
	AddressLine2  string
	City          string
	State         string
	ZIP           string
}
