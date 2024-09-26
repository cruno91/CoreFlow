-- Rollback: Drop content_entities table
BEGIN;
DROP TABLE IF EXISTS content_entities CASCADE;
COMMIT;