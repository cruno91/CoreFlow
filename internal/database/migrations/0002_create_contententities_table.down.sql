-- Rollback: Drop content_entities table
BEGIN;
-- Rollback: Drop content_entities table and all partitions
DROP TABLE IF EXISTS content_entities_hash_0 CASCADE;
DROP TABLE IF EXISTS content_entities_hash_1 CASCADE;
DROP TABLE IF EXISTS content_entities_hash_2 CASCADE;
DROP TABLE IF EXISTS content_entities_hash_3 CASCADE;
DROP TABLE IF EXISTS content_entities CASCADE;

COMMIT;