package liquor

import (
	"github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
)

var Licenses = aspect.Table("licenses",
	aspect.Column("id", aspect.String{PrimaryKey: true}),
	aspect.Column("bfn", aspect.String{}),
	aspect.Column("license", aspect.String{}),
	aspect.Column("name", aspect.String{}),
	aspect.Column("address", aspect.String{}),
	aspect.Column("code", aspect.String{}),
	aspect.Column("category", aspect.String{}),
	aspect.Column("license_name", aspect.String{}),
	aspect.Column("description", aspect.String{}),
	aspect.Column("issued", aspect.Timestamp{WithTimezone: true}),
	aspect.Column("expires", aspect.Timestamp{WithTimezone: true}),
	aspect.Column("status", aspect.String{}),
	aspect.Column("add_id", aspect.String{}),
	aspect.Column("external_address_id", aspect.String{}),
	aspect.Column("police_district", aspect.String{}),
	aspect.Column("council_district", aspect.String{}),
	aspect.Column("census_tract", aspect.String{}),
	aspect.Column("override", aspect.String{}),
	aspect.Column("longitude", aspect.Real{}),
	aspect.Column("latitude", aspect.Real{}),
	aspect.Column("location", postgis.Geometry{postgis.Point{}, 4326}),
)
