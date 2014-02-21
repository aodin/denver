Checkbook SQL
=============

http://data.denvergov.org/dataset/city-and-county-of-denver-checkbook

Schema:

    CREATE TABLE "checkbook" (
        "id" varchar PRIMARY KEY,
        "payee" varchar,
        "city" varchar,
        "state" varchar,
        "program_area" varchar,
        "department" varchar,
        "payment_id" varchar,
        "payment_date" timestamp,
        "year" int,
        "amount" real,
        "purchase_order" varchar,
        "funding_source" varchar,
        "project" varchar,
        "project_description" varchar,
        "expense_category" varchar,
        "expense_subcategory" varchar
    );

The only change made to the data was the removal of a double quote in the Project Description `"'Maintenance" RFID/Tech (temp)"` that exists on multiple lines.

The PostGres `COPY` command was then used to quickly load the data:

    COPY "checkbook" FROM 'checkbook.csv' WITH CSV HEADER;
