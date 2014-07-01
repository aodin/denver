package main

import (
	"fmt"
	"github.com/aodin/argo"
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/denver/crime"
	"github.com/aodin/denver/grocery"
	"github.com/aodin/denver/liquor"
	"github.com/aodin/volta/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{
    "version 1": "/v1/"
}`))
}

func main() {
	// Get config
	c, err := config.Parse()
	if err != nil {
		panic(err)
	}

	// Connect
	db, err := aspect.Connect(
		c.Database.Driver,
		c.Database.Credentials(),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := httprouter.New()

	router.GET("/", Root)

	baseURL := "/v1"
	api := argo.New(c, router, baseURL)
	api.Add("hearings", liquor.NewHearingsAPI(db))
	api.Add("licenses", liquor.NewLicensesAPI(db))
	api.Add("stores", grocery.NewStoresAPI(db))
	api.Add("offense-codes", crime.NewCodesAPI(db))

	address := fmt.Sprintf(":%d", c.Port)
	fmt.Printf("Starting on %s\n", address)
	if err = http.ListenAndServe(address, router); err != nil {
		panic(err)
	}
}
