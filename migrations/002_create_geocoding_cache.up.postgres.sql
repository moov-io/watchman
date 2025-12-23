-- Geocoding cache table for persistent storage of geocoding results.
-- This reduces API calls to geocoding providers by caching results.
CREATE TABLE geocoding_cache (
    cache_key  VARCHAR(512) NOT NULL PRIMARY KEY,
    latitude   DOUBLE PRECISION NOT NULL,
    longitude  DOUBLE PRECISION NOT NULL,
    accuracy   VARCHAR(20) NOT NULL DEFAULT 'unknown',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX geocoding_cache_created_at_idx ON geocoding_cache (created_at);
