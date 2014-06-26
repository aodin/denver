CREATE TABLE "grocery_stores" (
  "name" VARCHAR,
  "type" VARCHAR,
  "phone" VARCHAR,
  "hours" VARCHAR,
  "accepts_snap" BOOL,
  "sic" INTEGER,
  "naics" INTEGER,
  "sales_volume" INTEGER,
  "branch_status" VARCHAR,
  "employees" INTEGER,
  "franchise" VARCHAR,
  "square_footage" VARCHAR,
  "latitude" REAL,
  "longitude" REAL,
  "address" VARCHAR,
  "address_line1" VARCHAR,
  "address_line2" VARCHAR,
  "city" VARCHAR,
  "state" VARCHAR,
  "zip" VARCHAR
);

COPY "grocery_stores" FROM '/tmp/stores_2014-06-25.csv' WITH CSV HEADER;

ALTER TABLE "grocery_stores" ADD COLUMN "id" SERIAL PRIMARY KEY;

ALTER TABLE "grocery_stores" ADD COLUMN "location" Geography(Point);

UPDATE "grocery_stores" SET "location" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)::geography;

CREATE INDEX ON "grocery_stores" USING GIST ("location");

-- Final schema
CREATE TABLE "grocery_stores" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR,
  "type" VARCHAR,
  "phone" VARCHAR,
  "hours" VARCHAR,
  "accepts_snap" BOOL,
  "sic" INTEGER,
  "naics" INTEGER,
  "sales_volume" INTEGER,
  "branch_status" VARCHAR,
  "employees" INTEGER,
  "franchise" VARCHAR,
  "square_footage" VARCHAR,
  "latitude" REAL,
  "longitude" REAL,
  "address" VARCHAR,
  "address_line1" VARCHAR,
  "address_line2" VARCHAR,
  "city" VARCHAR,
  "state" VARCHAR,
  "zip" VARCHAR,
  "location" geometry(POINT, 4326)
);
