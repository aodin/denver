Addresses
=========

Addresses for the city of Denver.

[Available on data.denvergov.org](http://data.denvergov.org/dataset/city-and-county-of-denver-addresses)

2014-06-05: 269,143 addresses

SQL
---

The addresses can be loaded into a PostGres database using the schema included in this folder and by using the following `COPY` command for comma-separated values:

```sql
COPY "addresses" FROM 'addresses.csv' WITH CSV HEADER;
```


Queries
-------

Some example queries:

### Distinct Street Types

```sql
SELECT DISTINCT "street_type" FROM "addresses";
```

Using a `GROUP BY`:

```sql
SELECT "street_type" AS "type", COUNT("address") AS "count" FROM "addresses" GROUP BY "street_type" ORDER BY "count" DESC;
```

type | count  
:--- | ------:
St   | 157,366
Ave  |  53,811
Pl   |  11,824
Blvd |  11,548
Way  |  11,439
Ct   |   9,579
Dr   |   6,933
Pkwy |   2,333
     |   2,042
Rd   |   1,092
Cir  |     822
Ln   |     213
Cres |     129
Hwy  |      12


### Distinct Street Names

Does not count street type or directions.

```sql
SELECT COUNT(DISTINCT "street") FROM "addresses";
```

Returns: `891`
