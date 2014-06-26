Performance
===========

Testing of the Geography v. Geometry types for latitude and longitudes.

Use of `ST_DWithin` v. `ST_Distance`.

Use of indexes and output of `EXPLAIN ANALYZE`

Use the test schema:

```sql
CREATE TABLE "locations" (
    id VARCHAR PRIMARY KEY,
    latitude REAL,
    longitude REAL
);

CREATE TABLE "locations" (
    id VARCHAR PRIMARY KEY,
    latitude REAL,
    longitude REAL,
    geog GEOGRAPHY(POINT),
    geom GEOMETRY(POINT, 4326)
);
```

```sql
COPY "locations" FROM 'locations.csv' WITH CSV;
```

```sql
SELECT AddGeometryColumn('locations', 'geom', 4326, 'POINT', 2);

ALTER TABLE "locations" ADD COLUMN "geog" geography(Point); 
```

Update the `geog` and `geom` columns with `latitude` and `longitude`. Don't forget that `longitude` is the `x` coordinate!

```sql
UPDATE "locations" SET "geog" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)::geography;

UPDATE "locations" SET "geom" = ST_SetSRID(ST_MakePoint("longitude", "latitude"), 4326)::geometry;
```

Perform the distance calculation, using the Denver Capitol:

```
ST_SetSRID(ST_Point(-104.984722, 39.739167), 4326)::geography
```

```sql
SELECT count("id") FROM "locations" WHERE ST_DWithin("geog", ST_SetSRID(ST_Point(-104.984722, 39.739167), 4326)::geography, 2000);
```

Returns `682` results.

Looking at the `EXPLAIN ANALYZE` output:

```sql
EXPLAIN ANALYZE SELECT count("id") FROM "locations" WHERE ST_DWithin("geog", ST_SetSRID(ST_Point(-104.984722, 39.739167), 4326)::geography, 2000);
```

     Aggregate  (cost=646.35..646.36 rows=1 width=12) (actual time=8.327..8.327 rows=1 loops=1)
       ->  Seq Scan on locations  (cost=0.00..646.30 rows=21 width=12) (actual time=0.410..7.982 rows=682 loops=1)
             Filter: ((geog && '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography) AND ('0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography && _st_expand(geog, 2000::double precision)) AND _st_dwithin(geog, '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography, 2000::double precision, true))
             Rows Removed by Filter: 1435
     Total runtime: 8.393 ms


And for geometry:

```sql
EXPLAIN ANALYZE SELECT count("id") FROM "locations" WHERE ST_DWithin("geom":geography, ST_SetSRID(ST_Point(-104.984722, 39.739167), 4326)::geography, 2000);
```

     Aggregate  (cost=662.25..662.26 rows=1 width=12) (actual time=7.702..7.703 rows=1 loops=1)
       ->  Seq Scan on locations  (cost=0.00..662.18 rows=28 width=12) (actual time=0.414..7.530 rows=682 loops=1)
             Filter: (((geom)::geography && '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography) AND ('0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography && _st_expand((geom)::geography, 2000::double precision)) AND _st_dwithin((geom)::geography, '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography, 2000::double precision, true))
             Rows Removed by Filter: 1435
     Total runtime: 7.763 ms

Now add indexes:

```sql
CREATE INDEX ON "locations" USING GIST ("geog");

CREATE INDEX ON "locations" USING GIST ("geom");
```

Repeat the queries:

For geography:

     Aggregate  (cost=182.93..182.94 rows=1 width=12) (actual time=7.214..7.214 rows=1 loops=1)
       ->  Bitmap Heap Scan on locations  (cost=18.62..182.88 rows=21 width=12) (actual time=1.113..6.880 rows=682 loops=1)
             Recheck Cond: (geog && '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography)
             Filter: (('0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography && _st_expand(geog, 2000::double precision)) AND _st_dwithin(geog, '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography, 2000::double precision, true))
             Rows Removed by Filter: 133
             ->  Bitmap Index Scan on locations_geog_idx  (cost=0.00..18.61 rows=315 width=0) (actual time=0.729..0.729 rows=815 loops=1)
                   Index Cond: (geog && '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography)
     Total runtime: 7.323 ms

And geometry:

     Aggregate  (cost=662.25..662.26 rows=1 width=12) (actual time=15.295..15.295 rows=1 loops=1)
       ->  Seq Scan on locations  (cost=0.00..662.18 rows=28 width=12) (actual time=0.425..14.941 rows=682 loops=1)
             Filter: (((geom)::geography && '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography) AND ('0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography && _st_expand((geom)::geography, 2000::double precision)) AND _st_dwithin((geom)::geography, '0101000020E6100000B6696CAF053F5AC027A435069DDE4340'::geography, 2000::double precision, true))
             Rows Removed by Filter: 1435
     Total runtime: 15.353 ms


### References

* [What are the pros and cons of PostGIS geography and geometry types?](http://gis.stackexchange.com/q/6681)
* [ST_Distance](http://postgis.refractions.net/docs/ST_Distance.html)
* [ST_DWithin](http://postgis.net/docs/ST_DWithin.html)
* [Geography](http://workshops.boundlessgeo.com/postgis-intro/geography.html)
