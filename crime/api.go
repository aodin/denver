package crime

import (
	"github.com/aodin/argo"
	"github.com/aodin/argo/relational"
	"github.com/aodin/aspect"
	"net/url"
)

type CodesAPI struct {
	db       *aspect.DB
	resource *relational.Resource
}

func (api *CodesAPI) Get(parameters url.Values) argo.Response {
	// Generate a SELECT statement from the given URL parameters
	stmt, meta := api.resource.Select(parameters)
	// TODO Response on error?
	// return argo.Response{
	//     ContentType: "application/json",
	//     StatusCode:  400,
	//     Message:     map[string]string{"error": err.Error()},
	// }

	codes := make([]Code, 0)
	// TODO error is ignored
	api.db.QueryAll(stmt, &codes)

	results := map[string]interface{}{
		"results": codes,
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

func NewCodesAPI(db *aspect.DB) *CodesAPI {
	return &CodesAPI{
		db:       db,
		resource: relational.ResourceFromTable(Codes),
	}
}
