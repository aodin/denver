Convert to PostGIS
==================

1. Download a `*.shp` file

2. Use the `shp2pgsql` program:

        shp2pgsql input.shp > output.sql

    Note: the output program loves `MULTIPOLYGON`s, we can adjust this later

3. Load into postgres

        sudo -u postgres psql -d database < output.sql


Various Changes
---------------

### Set the SRID

    SELECT UpdateGeometrySRID('table', 'column', 4326);


### Convert to a Single Polygon

    ALTER TABLE "table" ADD COLUMN "polygon" geometry('POLYGON', 4326);
    UPDATE "table" SET "polygon" = ST_GeometryN("geom", 1);


### Create an Index

    CREATE INDEX ON "table" USING GIST ("column");


### Output as GeoJSON

    SELECT ST_AsGeoJSON("polygon") FROM "table";
