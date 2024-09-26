-- Migration: Create config_entities table
BEGIN;
CREATE TABLE IF NOT EXISTS config_entities (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    fields JSONB NOT NULL,
    path_pattern TEXT,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes for efficient querying (optional but recommended)
CREATE INDEX IF NOT EXISTS idx_config_entities_name ON config_entities(name);
COMMIT;