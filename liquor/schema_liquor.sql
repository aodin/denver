CREATE TABLE "licenses" (
    "id" varchar PRIMARY KEY,
    "bfn" varchar,
    "license" varchar,
    "name" varchar,
    "address" varchar,
    "code" varchar,
    "category" varchar,
    "license_name" varchar,
    "description" varchar,
    "issued" timestamp,
    "expires" timestamp,
    "status" varchar,
    "add_id" varchar,
    "external_address_id" varchar,
    "police_district" varchar,
    "council_district" varchar,
    "census_tract" varchar,
    "override" varchar,
    "longitude" real,
    "latitude" real
);

ALTER TABLE "licenses" ADD COLUMN "location" geography(Point);

UPDATE "licenses" SET "location" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)::geography;

CREATE INDEX ON "licenses" USING GIST ("location");