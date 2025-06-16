CREATE TABLE ingested_entities(
        type      VARCHAR(20) NOT NULL,
        source    VARCHAR(30) NOT NULL,
        source_id VARCHAR(40) NOT NULL,
        entity    JSONB NOT NULL
);

CREATE UNIQUE INDEX ingested_entities_uniq_idx ON ingested_entities (source, source_id);

CREATE INDEX ingested_entities_source_idx on ingested_entities (source, source_id ASC);
