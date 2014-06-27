package grocery

import (
	"github.com/aodin/argo"
	"github.com/aodin/argo/relational"
	"github.com/aodin/aspect"
	"net/url"
)

type StoresAPI struct {
	db       *aspect.DB
	resource *relational.Resource
}

func (api *StoresAPI) Get(parameters url.Values) argo.Response {
	// Generate a SELECT statement from the given URL parameters
	stmt, meta := api.resource.Select(parameters)
	// TODO Response on error?
	// return argo.Response{
	//     ContentType: "application/json",
	//     StatusCode:  400,
	//     Message:     map[string]string{"error": err.Error()},
	// }

	stores := make([]StoreWithId, 0)
	// TODO error is ignored
	api.db.QueryAll(stmt, &stores)

	results := map[string]interface{}{
		"results": stores,
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

func NewStoresAPI(db *aspect.DB) *StoresAPI {
	return &StoresAPI{
		db:       db,
		resource: relational.ResourceFromTable(Stores),
	}
}
