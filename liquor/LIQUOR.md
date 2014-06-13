Liquor
======

### Hearings

Public hearings are listed on [denvergov.org](http://www.denvergov.org/businesslicensing/DenverBusinessLicensingCenter/PublicHearingSchedule/tabid/441585/Default.aspx), which contains an inner frame with the [Hearing Viewer Application](http://www.denvergov.org/HearingViewerApplication/default.aspx).

### Licenses

A full list of liquor licenses is available from the City of Denver on [data.denvergov.org](http://data.denvergov.org/dataset/city-and-county-of-denver-liquor-licenses).

It is updated daily.

Each row in the file consists of 20 fields:

1. UNIQUE_ID
2. BFN
3. LIC_ID
4. BUS_PROF_NAME
5. FULL_ADDRESS
6. CODE
7. CATEGORY
8. LIC_NAME
9. DESCRIPTION
10. IDATE
11. EDATE
12. LIC_STATUS
13. ADD_ID
14. EXTERNAL_ADDRESS_ID
15. POLICE_DIST
16. COUNCIL_DIST
17. CENSUS_TRACT
18. OVERRIDE
19. X_COORD
20. Y_COORD

Fields 19 and 20 are a location on the [Colorado Central Grid](http://spatialreference.org/ref/epsg/2232/), also known as `EPSG:2232: NAD83 / Colorado Central (ftUS)`.

They can be converted to latitude and longitude using `gdaltransform`:

    gdaltransform -s_srs EPSG:2232 -t_srs EPSG:4326

This command can be installed on Debian/Ubuntu with the package `gdal-bin`.

Or as a Go command:

```go
coords := make([]string, len(licenses))
for i, license := range licenses {
    coords[i] = fmt.Sprintf("%s %s", license[18], license[19])
}

cmd := exec.Command(
    "gdaltransform",
    "-s_srs",
    "EPSG:2232",
    "-t_srs",
    "EPSG:4326",
)

// Create a single argument string
cmd.Stdin = strings.NewReader(strings.Join(coords, "\n"))

// Return as a bytes buffer
var out bytes.Buffer
cmd.Stdout = &out
err = cmd.Run()
if err != nil {
    return err
}

// The command returns as an x, y, and z coordinate
xyzs := strings.Split(out.String(), "\n")

var lnglat []string
for i, license := range licenses {
    // Remember the longitude is an x coordinate!
    lnglat = strings.SplitN(xyzs[i], " ", 3)
    license[18] = lnglat[0]
    license[19] = lnglat[1]
}
```

More information on the state plane coordinate system:

* [Wikipedia](http://en.wikipedia.org/wiki/State_Plane_Coordinate_System)
* [Earthpoint](http://www.earthpoint.us/StatePlane.aspx)


### Bounding Box

    <westbc>-105.109336</westbc>
    <eastbc>-104.671208</eastbc>
    <northbc>39.863214</northbc>
    <southbc>39.614990</southbc>

### Categories

As of 2014-02-01:

* LI32 - LIQUOR-3.2 % BEER          178
* LIRE - LIQUOR-RETAIL              204
* LIBW - LIQUOR-BEER & WINE          88
* TAST - LIQUOR TASTING              33
* LITA - LIQUOR-TAVERN              254
* LICL - LIQUOR-CLUB                 25
* LIAR - LIQUOR-ARTS                 10
* CAB8 - UNDERAGE CABARET PATRON     21
* LIHR - LIQUOR-HOTEL/RESTAURANT    817
* CABA - CABARET                    401
* LIDR - LIQUOR-DRUG STORE            3
* LIBR - LIQUOR-BREW-PUB              8
