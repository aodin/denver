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