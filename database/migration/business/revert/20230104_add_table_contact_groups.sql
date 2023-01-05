-- Revert business:20230104_add_table_contact_groups from pg

BEGIN;

DROP TABLE IF EXISTS "business"."contact_groups";

COMMIT;
