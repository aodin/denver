CREATE TABLE "crimes" (
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