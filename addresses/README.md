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

Latitude and Longitude
----------------------

Longitude (the x coordinate) varies greatly across the globe, ranging from 111.321 km at the equator to zero at the poles.

Latitude varies only slightly; from 110.567 km at the equator to 111.699 km at the poles.

The captiol building of Denver is at 39.739167, -104.984722

At the latitude and longitude, the difference between digits in both coordinates implies:

Using http://www.csgnetwork.com/degreelenllavcalc.html

### Latitude in Denver

Degrees | meters
------: | ----------:
1.0000  | 111,029.597
0.1000  |  11,102.960
0.0100  |   1,110.296
0.0010  |     111.030
0.0001  |      11.103


### Longitude in Denver

Degrees | meters
------: | ----------:
1.0000  |  85,717.848
0.1000  |   8,571.785
0.0100  |     857.178
0.0010  |      85.718
0.0001  |       8.572


Queries
-------

Some example queries:

### Distinct Street Types

```sql
SELECT DISTINCT "street_type" FROM "addresses";
```

Using a `GROUP BY`:

```sql
SELECT
    "street_type" AS "type",
    COUNT("address") AS "count"
FROM
    "addresses"
GROUP BY
    "street_type"
ORDER BY
    "count" DESC;
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

Output: `891`

Including street type:

```sql
SELECT COUNT (DISTINCT concat_ws(' ', street, street_type)) from "addresses";
```

Output: `1316`

With directions:

```
SELECT COUNT (DISTINCT concat_ws(' ', predirectional, street, street_type, postdirectional)) from "addresses";
```

Output: `1665`


### Duplicate Addresses


Duplicate full addresses; these are likely errors

```sql
SELECT 
    "address",
    "c"
FROM 
    (
        SELECT
            "address",
            COUNT("address") AS "c"
        FROM
            "addresses"
        GROUP BY
            "address"
    ) AS "counts"
WHERE
    "c" > 1;
```

Duplicate number and street

```sql
SELECT COUNT (DISTINCT concat_ws(' ', number, predirectional, street, street_type, postdirectional)) from "addresses";
```

Output: `189954`

### Most Common Addresses


```sql
SELECT 
    "unitless",
    "c"
FROM 
    (
        SELECT
            concat_ws(
                ' ',
                "number",
                "predirectional",
                "street",
                "street_type",
                "postdirectional"
            ) AS "unitless",
            COUNT("address") AS "c"
        FROM
            "addresses"
        GROUP BY
            "unitless"
    ) AS "counts"
WHERE
    "c" > 1
ORDER BY
    "c" DESC;
```

             address               |  c  
:--------------------------------- | --:
1020 15th St                       | 843
4760 S Wadsworth Blvd              | 658
601 W 11th Ave                     | 560
891 14th St                        | 508
8110 E Union Ave                   | 503
4380 S Monaco St                   | 458
10150 E Virginia Ave               | 417
3500 Rockmont Dr                   | 407
1700 Bassett St                    | 389
2905 N Inca St                     | 384
1164 S Acoma St                    | 373


```sql
SELECT * FROM "addresses" WHERE concat_ws(' ', "number", "predirectional", "street", "street_type", "postdirectional") = '2905.0 N Inca St';
```


