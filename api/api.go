package api

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/aspect"
	"github.com/aodin/denver/liquor"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
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
	var hearings []liquor.Hearing
	stmt := liquor.Hearings.Select().OrderBy(liquor.Hearings.C["id"].Asc())
	// TODO error is ignored
	h.db.QueryAll(stmt, &hearings)
	b, _ := json.MarshalIndent(&hearings, "", "    ")
	return b
}

type API struct {
	version   int
	db        *aspect.DB
	resources Resources
	router    *httprouter.Router
}

// Resources lists the available resources
func (a *API) Resources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	root := fmt.Sprintf("v%d", a.version)
	b, _ := json.MarshalIndent(a.resources.URLs(root), "", "    ")
	w.Write(b)
}

func (a *API) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Determine which resource was requested
	name := ps.ByName("resource")
	resource, exists := a.resources[name]
	if !exists {
		http.NotFound(w, r)
		return
	}
	// TODO Errors are ignored
	// Get the parameters
	// TODO Or r.URL.Query()
	r.ParseForm()
	w.Write(resource.Get(r.Form))
}

func (a *API) Add(name string, resource Resource) error {
	if _, exists := a.resources[name]; exists {
		return fmt.Errorf("Resource %s already exists", name)
	}
	a.resources[name] = resource
	return nil
}

func (a *API) ListenAndServe() error {
	// TODO Config
	fmt.Println("Starting on :8080")
	return http.ListenAndServe(":8080", a.router)
}

// TODO Root should include all attached APIs
func Root(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte(fmt.Sprintf(`{"version %d": "/v%d/"}`, 1, 1)))
}

func New(db *aspect.DB) *API {
	a := &API{
		version:   1,
		db:        db,
		resources: make(Resources),
		router:    httprouter.New(),
	}
	a.router.GET("/", Root)
	a.router.GET("/v1/", a.Resources)
	a.Add("hearings", &hearingsAPI{db: db})
	a.router.GET("/v1/:resource", a.Get)
	return a
}
