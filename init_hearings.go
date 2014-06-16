package main

import (
	"flag"
	"fmt"
	"github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/denver/config"
	"github.com/aodin/denver/liquor"
	"io/ioutil"
	"os"
)

func InitSQL() {
	// TODO Actually create the table
	fmt.Println(liquor.Hearings.Create())
}

// TODO Does not clear previous entries
func LoadFile(db *aspect.DB, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	hearings, err := liquor.ParseHearingsJSON(contents)
	if err != nil {
		return err
	}
	stmt := liquor.Hearings.Insert(hearings)
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
	config, err := config.ParseDatabase()
	if err != nil {
		panic(err)
	}

	if init {
		InitSQL()
	}

	if load != "" {
		// Connect
		db, err := aspect.Connect(
			config.Driver,
			config.Credentials(),
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
