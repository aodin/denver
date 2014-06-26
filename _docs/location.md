Location
========

How to add location to a PostGIS table.


1. Add a `Geography` column.


```sql
ALTER TABLE "licenses" ADD COLUMN "location" Geography(Point); 
```

2. Update the `Geography` column with `latitude` and `longitude`. Don't forget that `longitude` is the `x` coordinate!

```sql
UPDATE "licenses" SET "location" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)::geography;
```

3. Add an index.

```sql
CREATE INDEX ON "licenses" USING GIST ("location");
```


### Example Query

```sql
SELECT "name", "address" FROM "licenses" WHERE ST_DWithin("location", ST_SetSRID(ST_Point(-104.984722, 39.739167), 4326)::geography, 200);
```