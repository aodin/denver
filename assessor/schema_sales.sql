CREATE TABLE "sales" (
    "schednum" varchar,
    "reception_num" varchar,
    "instrument" varchar,
    "year" int,
    "monthday" int,
    "reception" timestamp,
    "price" varchar,
    "grantor" varchar,
    "grantee" varchar,
    "class" varchar,
    "mkt_clus" varchar,
    "d_class" varchar,
    "d_class_cn" varchar,
    "nbhd_id" int,
    "nbhd_name" varchar,
    "pin" varchar
);

CREATE INDEX ON "sales" ("nbhd_id");

UPDATE "sales" SET "price" = NULL WHERE "price" = '';
UPDATE "sales" SET "pin" = NULL WHERE "pin" = '';

ALTER TABLE "sales" ALTER COLUMN "price" TYPE real USING ("price"::real);
ALTER TABLE "sales" ALTER COLUMN "pin" TYPE real USING ("pin"::int);

-- Some fields need replacement because of improper escaping

-- "O"CONNOR,MATTHEW B", "O'CONNOR,MATTHEW B"
-- "GAMBILL,SHANONNON O" QUINN &", "GAMBILL,SHANONNON O' QUINN &"

