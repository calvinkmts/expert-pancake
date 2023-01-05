-- Revert business:20230504_add_table_contact_groups_and_contact_books from pg

BEGIN;

DROP TABLE IF EXISTS "business"."contact_groups";
DROP TABLE IF EXISTS "business"."contact_books";

COMMIT;
