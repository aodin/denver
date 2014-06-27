package liquor

import (
	"github.com/aodin/argo"
	"github.com/aodin/argo/relational"
	"github.com/aodin/aspect"
	"net/url"
)

type HearingsAPI struct {
	db       *aspect.DB
	resource *relational.Resource
}

func (api *HearingsAPI) Get(parameters url.Values) argo.Response {
	// Generate a SELECT statement from the given URL parameters
	stmt, meta := api.resource.Select(parameters)
	// TODO Response on error?
	// return argo.Response{
	//     ContentType: "application/json",
	//     StatusCode:  400,
	//     Message:     map[string]string{"error": err.Error()},
	// }

	hearings := make([]Hearing, 0)
	// TODO error is ignored
	api.db.QueryAll(stmt, &hearings)

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
	return &HearingsAPI{
		db:       db,
		resource: relational.ResourceFromTable(Hearings),
	}
}

type LicensesAPI struct {
	db       *aspect.DB
	resource *relational.Resource
}

// Example lat/long query:
// /v1/licenses/?latitude=39.739167&longitude=-104.984722
func (api *LicensesAPI) Get(parameters url.Values) argo.Response {
	// Generate a SELECT statement from the given URL parameters
	stmt, meta := api.resource.Select(parameters)
	// TODO Response on error?
	// return argo.Response{
	//     ContentType: "application/json",
	//     StatusCode:  400,
	//     Message:     map[string]string{"error": err.Error()},
	// }

	// Prevent a "null" response which can happen with var licenses []License
	licenses := make([]License, 0)

	// TODO error is ignored
	api.db.QueryAll(stmt, &licenses)

	results := map[string]interface{}{
		"results": licenses,
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

func NewLicensesAPI(db *aspect.DB) *LicensesAPI {
	return &LicensesAPI{
		db:       db,
		resource: relational.ResourceFromTable(Licenses),
	}
}
