CREATE SCHEMA IF NOT EXISTS business;

CREATE TABLE business.companies (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "name" text NOT NULL,
  "initial_name" text NOT NULL,
  "image_url" text NOT NULL,
  "type" text NOT NULL,
  "responsible_person" text NOT NULL,
  "is_deleted" boolean DEFAULT FALSE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE business.company_branches (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "address" text NOT NULL,
  "phone_number" text NOT NULL,
  "is_central" boolean DEFAULT FALSE NOT NULL,
  "is_deleted" boolean DEFAULT FALSE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "business"."company_members" (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "business"."company_member_requests" (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "status" text NOT NULL DEFAULT 'waiting',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);