package main

import (
	"flag"
	"fmt"
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/denver/grocery"
	"github.com/aodin/volta/config"
)

func InitSQL() {
	// TODO Actually create the table
	fmt.Println(grocery.Stores.Create())
}

// TODO Does not clear previous entries
func LoadFile(db *aspect.DB, path string) error {
	stores, err := grocery.ParseConvertedStoresCSV(path)
	if err != nil {
		return err
	}
	stmt := grocery.Stores.Insert(stores)
	_, err = db.Execute(stmt)
	return err
}

func main() {
	var init bool
	var load string
	flag.BoolVar(&init, "init", false, "print SQL for CREATE TABLE")
	flag.StringVar(&load, "load", "", "load the given file")
	flag.Parse()

	// Get the database driver
	c, err := config.ParseFile("../../settings.json")
	if err != nil {
		panic(err)
	}

	if init {
		InitSQL()
	}

	if load != "" {
		// Connect
		db, err := aspect.Connect(
			c.Database.Driver,
			c.Database.Credentials(),
		)
		if err != nil {
			panic(err)
		}
		if err = LoadFile(db, load); err != nil {
			panic(err)
		}
	}
}

// CREATE TABLE "hearings" (
//   "id" INTEGER PRIMARY KEY,
//   "notice_link" VARCHAR,
//   "name" VARCHAR,
//   "address" VARCHAR,
//   "latitude" REAL,
//   "longitude" REAL,
//   "time" TIMESTAMP WITH TIME ZONE,
//   "outcome" VARCHAR
// );
