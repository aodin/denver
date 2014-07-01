package crime

import (
	"time"
)

type rawCrime struct {
	IncidentID      int64
	OffenseID       int64
	Code            int64
	CodeExt         int64
	Type            string
	Category        string
	FirstOccurrence string
	LastOccurrence  string
	Reported        string
	Address         string
	X               int64
	Y               int64
	Latitude        float64
	Longitude       float64
	District        int64
	Precinct        int64
	Neighborhood    string
}

type Crime struct {
	IncidentID      int64      `json:"incident_id"`
	OffenseID       int64      `json:""`
	CodeID          string     `json:"code_id"`
	Code            int64      `json:"code"`
	CodeExt         int64      `json:"code_extension"`
	Type            string     `json:"type"`
	Category        string     `json:"category"`
	FirstOccurrence time.Time  `json:"first_occurrence"`
	LastOccurrence  *time.Time `json:"last_occurrence"`
	Reported        time.Time  `json:"reported"`
	Address         string     `json:"address"`
	Latitude        float64    `json:"latitude"`
	Longitude       float64    `json:"longitude"`
	District        int64      `json:"district"`
	Precinct        int64      `json:"precinct"`
	Neighborhood    string     `json:"neighborhood"`
}
