package api

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/aspect"
	"github.com/aodin/denver/config"
	"github.com/aodin/denver/liquor"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
	"strings"
)

// TODO headers too?
type Resource interface {
	Get(url.Values) []byte
}

type Resources map[string]Resource

func (r Resources) URLs(root string) map[string]string {
	urls := make(map[string]string)
	for key, _ := range r {
		urls[key] = fmt.Sprintf("/%s/%s/", root, key)
	}
	return urls
}

type hearingsAPI struct {
	db *aspect.DB
}

func (h *hearingsAPI) Get(parameters url.Values) []byte {
	// Allow ordering from GET parameters
	order := strings.ToLower(parameters.Get("order"))

	// We also need to remove the minus sign if it exists
	var inverted bool
	if order != "" && string(order[0]) == `-` {
		order = order[1:]
		inverted = true
	}

	// Default to "id" ASC
	_, exists := liquor.Hearings.C[order]
	if order == "" || !exists {
		order = "id"
		inverted = false
	}

	// Create the order by statements
	orderBy := liquor.Hearings.C[order].Asc()
	if inverted {
		orderBy = orderBy.Desc()
	}

	// Perform the query and return all results
	// TODO pagination
	stmt := liquor.Hearings.Select().OrderBy(orderBy)
	var hearings []liquor.Hearing
	// TODO error is ignored
	h.db.QueryAll(stmt, &hearings)

	// Convert to pretty printed JSON
	b, _ := json.MarshalIndent(&hearings, "", "    ")
	return b
}

type API struct {
	version   int
	db        *aspect.DB
	resources Resources
	router    *httprouter.Router
	config    config.Config
}

// Resources lists the available resources
func (a *API) Resources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	root := fmt.Sprintf("v%d", a.version)
	b, _ := json.MarshalIndent(a.resources.URLs(root), "", "    ")
	w.Write(b)
}

// Get will GET the resource with the requested name
func (a *API) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Determine which resource was requested
	name := ps.ByName("resource")
	resource, exists := a.resources[name]
	if !exists {
		http.NotFound(w, r)
		return
	}
	// Set the content type
	w.Header().Set("Content-Type", "application/json")

	// Get the parameters
	w.Write(resource.Get(r.URL.Query()))
}

// Add will add the given resource at the given name
func (a *API) Add(name string, resource Resource) error {
	if _, exists := a.resources[name]; exists {
		return fmt.Errorf("Resource %s already exists", name)
	}
	a.resources[name] = resource
	return nil
}

// ListenAndServe will start the API server and run forever
func (a *API) ListenAndServe() error {
	address := fmt.Sprintf(":%d", a.config.Port)
	fmt.Printf("Starting on %s\n", address)
	return http.ListenAndServe(address, a.router)
}

// TODO Root should include all attached APIs
func Root(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{
    "version %d": "/v%d/"
}`, 1, 1)))
}

// New will create a new API instance
// TODO A global registry for API versions
func New(c config.Config, db *aspect.DB) *API {
	a := &API{
		version:   1,
		db:        db,
		resources: make(Resources),
		router:    httprouter.New(),
		config:    c,
	}
	a.router.GET("/", Root)
	a.router.GET("/v1/", a.Resources)
	a.Add("hearings", &hearingsAPI{db: db})
	// TODO This should be added automatically
	a.router.GET("/v1/:resource/", a.Get)
	return a
}
