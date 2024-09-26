-- Migration: Create content_entities table
BEGIN;
CREATE TABLE IF NOT EXISTS content_entities (
    id UUID PRIMARY KEY,
    config_entity_id UUID NOT NULL REFERENCES config_entities(id) ON DELETE CASCADE,
    fields JSONB NOT NULL,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes for efficient querying
CREATE INDEX IF NOT EXISTS idx_content_entities_config_entity_id ON content_entities(config_entity_id);
CREATE INDEX IF NOT EXISTS idx_content_entities_created_at ON content_entities(created_at);
COMMIT;