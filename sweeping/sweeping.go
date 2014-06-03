package main

import (
	"encoding/json"
	"fmt"
	. "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
	"github.com/kpawlik/geojson"
)

var Schedules = Table("street_sweep_schedule",
	Column("gid", Integer{PrimaryKey: true}),
	Column("months", String{}),
	Column("week", String{}),
	Column("days", String{}),
	Column("posted", String{}),
	Column("sssid", String{}),
	Column("polygon", postgis.Geometry{postgis.Polygon{}, 4326}),
)

// TODO Let's cheat with the JSONMarshaler interface
type Schedule struct {
	Id      int64           `json:"gid"   db:"id"`
	Months  string          `json:"months" db:"months"`
	Week    string          `json:"week" db:"week"`
	Days    string          `json:"days" db:"days"`
	Posted  string          `json:"posted" db:"posted"`
	Polygon json.RawMessage `json:"polygon" db:"polygon"`
}

type Fixed struct {
	Id       int64           `json:"id"`
	Months   string          `json:"months"`
	Week     string          `json:"week"`
	Days     string          `json:"days"`
	Posted   string          `json:"posted"`
	Geometry geojson.Polygon `json:"geometry"`
}

func Convert(s Schedule) (f Fixed) {
	f.Id = s.Id
	f.Months = s.Months
	f.Week = s.Week
	f.Days = s.Days
	f.Posted = s.Posted
	json.Unmarshal(s.Polygon, &f.Geometry)
	return
}

func main() {
	db, err := Connect(
		"postgres",
		"host=localhost port=5432 dbname=denver user=postgres password=gotest",
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt := Select(
		Schedules.C["gid"],
		Schedules.C["months"],
		Schedules.C["week"],
		Schedules.C["days"],
		Schedules.C["posted"],
		postgis.AsGeoJSON(Schedules.C["polygon"]),
	).OrderBy(Schedules.C["gid"])

	// Get all neighborhoods
	var schedules []Schedule
	if err := db.QueryAll(stmt, &schedules); err != nil {
		panic(err)
	}

	output := make([]Fixed, len(schedules))
	for i, schedule := range schedules {
		output[i] = Convert(schedule)
	}

	b, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
