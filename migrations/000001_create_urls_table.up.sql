CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(10) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    tittle VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    click_count BIGINT NOT NULL DEFAULT 0,
    last_accessed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_urls_code ON urls(code);
CREATE INDEX idx_urls_is_active on urls(is_active);
CREATE INDEX idx_urls_expires_at on urls(expires_at) WHERE is_active IS NOT NULL;