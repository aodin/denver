package crime

import (
	"github.com/aodin/aspect"
)

var Codes = aspect.Table("offense_codes",
	aspect.Column("id", aspect.String{PrimaryKey: true}),
	aspect.Column("code", aspect.Integer{}),
	aspect.Column("extension", aspect.Integer{}),
	aspect.Column("description", aspect.String{}),
	aspect.Column("category", aspect.String{}),
	aspect.Column("is_crime", aspect.Boolean{}),
	aspect.Column("is_traffic", aspect.Boolean{}),
)
