package grocery

import (
	"github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
	"github.com/aodin/aspect/postgres"
)

var Stores = aspect.Table("grocery_stores",
	aspect.Column("id", postgres.Serial{PrimaryKey: true}),
	aspect.Column("name", aspect.String{}),
	aspect.Column("type", aspect.String{}),
	aspect.Column("phone", aspect.String{}),
	aspect.Column("hours", aspect.String{}),
	aspect.Column("accepts_snap", aspect.Boolean{}),
	aspect.Column("sic", aspect.Integer{}),
	aspect.Column("naics", aspect.Integer{}),
	aspect.Column("sales_volume", aspect.Integer{}),
	aspect.Column("branch_status", aspect.String{}),
	aspect.Column("employees", aspect.Integer{}),
	aspect.Column("franchise", aspect.String{}),
	aspect.Column("square_footage", aspect.String{}),
	aspect.Column("latitude", aspect.Real{}),
	aspect.Column("longitude", aspect.Real{}),
	aspect.Column("address", aspect.String{}),
	aspect.Column("address_line1", aspect.String{}),
	aspect.Column("address_line2", aspect.String{}),
	aspect.Column("city", aspect.String{}),
	aspect.Column("state", aspect.String{}),
	aspect.Column("zip", aspect.String{}),
	aspect.Column("location", postgis.Geometry{postgis.Point{}, 4326}),
)
