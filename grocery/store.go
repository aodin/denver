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

type StoreWithId struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Phone         string  `json:"phone"`
	Hours         string  `json:"hours"`
	AcceptsSnap   bool    `json:"accepts_snap"`
	SIC           int64   `json:"SIC"`
	NAICS         int64   `json:"NAICS"`
	SalesVolume   int64   `json:"sales_volume"`
	BranchStatus  string  `json:"branch_status"`
	Employees     int64   `json:"employees"`
	Franchise     string  `json:"franchise"`
	SquareFootage string  `json:"square_footage"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Address       string  `json:"address"`
	AddressLine1  string  `json:"address_line_1"`
	AddressLine2  string  `json:"address_line_2"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	ZIP           string  `json:"zip"`
}
