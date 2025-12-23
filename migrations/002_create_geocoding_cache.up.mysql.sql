-- Geocoding cache table for persistent storage of geocoding results.
-- This reduces API calls to geocoding providers by caching results.
CREATE TABLE geocoding_cache (
    cache_key  VARCHAR(64) NOT NULL,
    latitude   DOUBLE NOT NULL,
    longitude  DOUBLE NOT NULL,
    accuracy   VARCHAR(20) NOT NULL DEFAULT 'unknown',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (cache_key),
    INDEX geocoding_cache_created_at_idx (created_at)
);
