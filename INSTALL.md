INSTALL
=======

Create a `settings.json` file for your current install:

```json
{
    "port": 8008,
    "database": {
        "driver": "postgres",
        "host": "localhost",
        "port": 5432,
        "name": "denver",
        "user": "postgres",
        "password": "password"
    } 
}
```

### API

The `github.com/aodin/denver/api` package uses the following packages:

* [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)


Third-Party Package Installations
---------------------------------

The following third party packages have additional requirements.

### Gokogiri

[Gokogiri](https://github.com/moovweb/gokogiri) is used to scrap HTML. It requires the following packages on Ubuntu 12.04:

    sudo apt-get install libxml2-dev pkg-config


### PostGIS

    sudo apt-get install postgis

Enable the extension on the PostGres database:

    CREATE EXTENSION postgis;
