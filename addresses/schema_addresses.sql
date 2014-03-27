CREATE TABLE "addresses" (
    "type" varchar,
    "x" real,
    "y" real,
    "latitude" real,
    "longitude" real,
    "number_prefix" varchar,
    "number" varchar,
    "number_suffix" varchar,
    "premodifier" varchar,
    "predirectional" varchar,
    "street" varchar,
    "street_type" varchar,
    "postdirectional" varchar,
    "postmodifier" varchar,
    "building_type" varchar,
    "building_id" varchar,
    "unit_type" varchar,
    "unit_id" varchar,
    "composite_unit_type" varchar,
    "composite_unit_id" varchar,
    "address" varchar
);

-- CREATE INDEX ON "addresses" ((lower("street")));
-- CREATE INDEX ON "addresses" ((lower("number")));
-- CREATE INDEX ON "addresses" ((lower("address")));

CREATE EXTENSION "pg_trgm";

CREATE INDEX ON "addresses" USING gin("address" gin_trgm_ops);
