-- Deploy business:20230104_add_table_contact_groups to pg

BEGIN;

CREATE TABLE "business"."contact_groups" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
