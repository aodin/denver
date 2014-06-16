package main

import (
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/denver/api"
	"github.com/aodin/denver/config"
)

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

	a := api.New(c, db)
	if err = a.ListenAndServe(); err != nil {
		panic(err)
	}
}
