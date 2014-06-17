package liquor

import (
	"github.com/aodin/argo"
	"github.com/aodin/aspect"
	"net/url"
	"strings"
)

type HearingsAPI struct {
	db *aspect.DB
}

func (h *HearingsAPI) Get(parameters url.Values) argo.Response {
	// Allow ordering from GET parameters
	order := strings.ToLower(parameters.Get("order"))

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
	stmt := Hearings.Select().OrderBy(orderBy)
	var hearings []Hearing
	// TODO error is ignored
	h.db.QueryAll(stmt, &hearings)

	return argo.Response{
		ContentType: "application/json",
		StatusCode:  200,
		Results:     hearings,
	}
}

func NewHearingsAPI(db *aspect.DB) *HearingsAPI {
	return &HearingsAPI{db}
}
