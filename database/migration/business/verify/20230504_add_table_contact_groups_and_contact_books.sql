-- Verify business:20230504_add_table_contact_groups_and_contact_books on pg

BEGIN;

DO $$ << if_contact_groups_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'contact_groups'
) THEN RAISE EXCEPTION 'table business.contact_groups not found';
END IF;
END if_contact_groups_table_exist_test $$;

DO $$ << if_contact_books_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'contact_books'
) THEN RAISE EXCEPTION 'table business.contact_books not found';
END IF;
END if_contact_books_table_exist_test $$;

ROLLBACK;
