-- Deploy business:20230504_add_table_contact_groups_and_contact_books to pg

BEGIN;

CREATE TABLE "business"."contact_groups" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "business"."contact_books" (
  "id" text NOT NULL,
  "primary_company_id" text NOT NULL,
  "secondary_company_id" text NOT NULL,
  "contact_group_id" text NOT NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NOT NULL,
  "mobile" text NOT NULL,
  "web" text NOT NULL,
  "mail_address" text NOT NULL,
  "shipping_address" text NOT NULL,
  "is_customer" bool NOT NULL DEFAULT false,
  "customer_pic" text NOT NULL,
  "is_supplier" bool NOT NULL DEFAULT false,
  "supplier_pic" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
