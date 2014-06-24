package liquor

import (
	"errors"
	"github.com/aodin/argo"
	"github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
	"net/url"
	"strconv"
	"strings"
)

var LatLngUnparsed = errors.New("liquor: no attempt was made to parse latlng")

type HearingsAPI struct {
	db *aspect.DB
}

func (h *HearingsAPI) Get(parameters url.Values) argo.Response {
	// Allow ordering from GET parameters
	order := strings.ToLower(parameters.Get("order"))

	radius := ParseInt(parameters.Get("radius"), 200)
	// TODO Error if radius is too large?

	// We also need to remove the minus sign if it exists
	var inverted bool
	if order != "" && string(order[0]) == `-` {
		order = order[1:]
		inverted = true
	}

	// Default to "id" ASC
	_, exists := Hearings.C[order]
	if order == "" || !exists {
		order = "id"
		inverted = false
	}

	// Create the order by statements
	orderBy := Hearings.C[order].Asc()
	if inverted {
		orderBy = orderBy.Desc()
	}

	// Perform the query and return all results
	// TODO pagination
	stmt := Hearings.SelectExcept(Hearings.C["location"]).OrderBy(orderBy)

	meta := map[string]interface{}{}

	lat, lng, err := ParseLatLng(parameters)
	if err == LatLngUnparsed {
		// Do nothing
	} else if err != nil {
		return argo.Response{
			ContentType: "application/json",
			StatusCode:  400,
			Message:     map[string]string{"error": err.Error()},
		}
	} else {
		// Lat and Lng were parsed successfully
		// Create a where statement
		stmt = stmt.Where(postgis.DWithin(
			Hearings.C["location"],
			postgis.LatLong{lat, lng},
			radius,
		))
		meta["latitude"] = lat
		meta["longitude"] = lng
		meta["radius"] = radius
	}

	var hearings []Hearing
	// TODO error is ignored
	h.db.QueryAll(stmt, &hearings)

	results := map[string]interface{}{
		"results": hearings,
	}

	if len(meta) > 0 {
		results["meta"] = meta
	}

	return argo.Response{
		ContentType: "application/json",
		StatusCode:  200,
		Message:     results,
	}
}

func NewHearingsAPI(db *aspect.DB) *HearingsAPI {
	return &HearingsAPI{db}
}

type LicensesAPI struct {
	db *aspect.DB
}

func ParseInt(input string, fallback int) int {
	parsed, err := strconv.Atoi(input)
	if err != nil {
		return fallback
	}
	return parsed
}

// ParseLatLng will look for the GET parameters "latitude" and "longitude".
// If either of these parameters exist and are non-empty, they will attempt
// to parse them as latitude and longitude values, returning an error
// if the parameters could no be successfully parsed or were out of range.
func ParseLatLng(parameters url.Values) (lat, lng float64, err error) {
	rawLat := parameters.Get("latitude")
	rawLng := parameters.Get("longitude")
	if rawLat == "" && rawLng == "" {
		err = LatLngUnparsed
		return
	}
	if lat, err = strconv.ParseFloat(rawLat, 64); err != nil {
		return
	}
	lng, err = strconv.ParseFloat(rawLng, 64)
	// TODO Error if out of bounds
	return
}

// Example lat/long query:
// /v1/licenses/?latitude=39.739167&longitude=-104.984722
func (h *LicensesAPI) Get(parameters url.Values) argo.Response {
	// Allow ordering from GET parameters
	order := strings.ToLower(parameters.Get("order"))

	// Determine the offset and limit
	// TODO Limit is hard-coded but should be a configurable option
	limit := 100
	offset := ParseInt(parameters.Get("offset"), 0)
	radius := ParseInt(parameters.Get("radius"), 200)
	// TODO Error if radius is too large?

	// TODO Determine if specific field were requested

	// We also need to remove the minus sign if it exists
	var inverted bool
	if order != "" && string(order[0]) == `-` {
		order = order[1:]
		inverted = true
	}

	// Default to "id" ASC
	_, exists := Licenses.C[order]
	if order == "" || !exists {
		order = "id"
		inverted = false
	}

	// Create the order by statements
	orderBy := Licenses.C[order].Asc()
	if inverted {
		orderBy = orderBy.Desc()
	}

	// Perform the query and return all results
	stmt := Licenses.SelectExcept(Licenses.C["location"]).OrderBy(orderBy).Limit(limit).Offset(offset)

	// Start meta
	meta := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	lat, lng, err := ParseLatLng(parameters)
	if err == LatLngUnparsed {
		// Do nothing
	} else if err != nil {
		return argo.Response{
			ContentType: "application/json",
			StatusCode:  400,
			Message:     map[string]string{"error": err.Error()},
		}
	} else {
		// Lat and Lng were parsed successfully
		// Create a where statement
		stmt = stmt.Where(postgis.DWithin(
			Licenses.C["location"],
			postgis.LatLong{lat, lng},
			radius,
		))
		meta["latitude"] = lat
		meta["longitude"] = lng
		meta["radius"] = radius
	}

	var licenses []License
	// TODO error is ignored
	h.db.QueryAll(stmt, &licenses)

	results := map[string]interface{}{
		"results": licenses,
		"meta":    meta,
	}

	return argo.Response{
		ContentType: "application/json",
		StatusCode:  200,
		Message:     results,
	}
}

func NewLicensesAPI(db *aspect.DB) *LicensesAPI {
	return &LicensesAPI{db}
}
