CREATE TABLE "raw_crimes" (
    "id" varchar,
    "offense_id" varchar,
    "offense_code" varchar,
    "offense_code_ext" varchar,
    "offense_type" varchar,
    "offense_category_id" varchar,
    "first_occurrence" timestamp,
    "last_occurrence" timestamp,
    "reported" timestamp,
    "address" varchar,
    "x" real,
    "y" real,
    "longitude" real,
    "latitude" real,
    "district" varchar,
    "precinct" varchar,
    "neighborhood" varchar
);

CREATE TABLE "offense_codes" (
  "id" VARCHAR PRIMARY KEY,
  "code" INTEGER,
  "extension" INTEGER,
  "description" VARCHAR,
  "category" VARCHAR,
  "is_crime" BOOL,
  "is_traffic" BOOL
);

COPY "offense_codes" FROM '/tmp/codes_2014-07-01.csv' WITH CSV HEADER;
