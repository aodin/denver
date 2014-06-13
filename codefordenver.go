package main

import (
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/denver/api"
	"github.com/aodin/denver/config"
)

func main() {
	// Get config
	config, err := config.Parse()
	if err != nil {
		panic(err)
	}
	// Connect
	db, err := aspect.Connect(
		config.Driver,
		config.Credentials(),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := api.New(db)
	if err = a.ListenAndServe(); err != nil {
		panic(err)
	}
}
