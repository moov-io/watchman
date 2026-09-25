-- Per-source ingest snapshot metadata. Search compares this one-row checksum
-- against the in-memory corpus so it does not scan ingested_entities on every query.
CREATE TABLE ingest_sources (
    source       VARCHAR(30) NOT NULL,
    entity_count INT NOT NULL,
    checksum     CHAR(64) NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (source)
);
