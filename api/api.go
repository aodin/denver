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
	db        *aspect.DB
	resources map[string]Resource
	router    *httprouter.Router
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

func New(db *aspect.DB) *API {
	a := &API{
		db:        db,
		resources: make(map[string]Resource),
		router:    httprouter.New(),
	}
	a.Add("hearings", &hearingsAPI{db: db})
	a.router.GET("/api/:resource", a.Get)
	return a
}
