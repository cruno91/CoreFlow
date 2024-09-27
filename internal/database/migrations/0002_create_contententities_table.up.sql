-- Migration: Create content_entities table
BEGIN;
-- Migration: Create partitioned content_entities table with subpartitions

-- Step 1: Create the parent table partitioned by HASH (config_entity_id)
CREATE TABLE IF NOT EXISTS content_entities (
    id UUID PRIMARY KEY,
    config_entity_id UUID NOT NULL,
    bundle TEXT NOT NULL,
    name TEXT NOT NULL,
    fields JSONB NOT NULL,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (config_entity_id) REFERENCES config_entities(id) ON DELETE CASCADE
) PARTITION BY HASH (config_entity_id);

-- Step 2: Create first-level hash partitions
CREATE TABLE content_entities_hash_0 PARTITION OF content_entities
    FOR VALUES WITH (MODULUS 4, REMAINDER 0) PARTITION BY LIST (bundle);

CREATE TABLE content_entities_hash_1 PARTITION OF content_entities
    FOR VALUES WITH (MODULUS 4, REMAINDER 1) PARTITION BY LIST (bundle);

CREATE TABLE content_entities_hash_2 PARTITION OF content_entities
    FOR VALUES WITH (MODULUS 4, REMAINDER 2) PARTITION BY LIST (bundle);

CREATE TABLE content_entities_hash_3 PARTITION OF content_entities
    FOR VALUES WITH (MODULUS 4, REMAINDER 3) PARTITION BY LIST (bundle);

-- Step 3: Create indexes on first-level partitions (optional)
-- Index on config_entity_id (though partitioning is already on this column)
-- Create indexes on other columns if needed

-- Since subpartitions (second-level) will be created dynamically, indexes on them will be added in the application code

COMMIT;