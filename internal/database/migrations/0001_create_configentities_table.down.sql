-- Rollback: Drop config_entities table
BEGIN;
DROP TABLE IF EXISTS config_entities CASCADE;
COMMIT;
