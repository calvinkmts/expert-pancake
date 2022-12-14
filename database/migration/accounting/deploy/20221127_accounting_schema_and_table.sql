-- Deploy accounting:20221127_accounting_schema_and_table to pg

BEGIN;

CREATE SCHEMA IF NOT EXISTS accounting;

CREATE TABLE "accounting"."company_fiscal_years" (
  "company_id" text NOT NULL,
  "start_month" int NOT NULL DEFAULT 0,
  "start_year" int NOT NULL DEFAULT 0,
  "end_month" int NOT NULL DEFAULT 0,
  "end_year" int NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

CREATE TABLE "accounting"."company_chart_of_accounts" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "account_code" text NOT NULL,
  "account_name" text NOT NULL,
  "account_group" text NOT NULL,
  "bank_name" text NOT NULL,
  "bank_account_number" text NOT NULL,
  "bank_code" text NOT NULL,
  "opening_balance" float NOT NULL DEFAULT 0,
  "is_deleted" int NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
