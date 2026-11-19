CREATE TABLE list_refs(
        id            BIGSERIAL PRIMARY KEY,
        list_name     VARCHAR(64) NOT NULL,
        list_hash     VARCHAR(64) NOT NULL,
        downloaded_at TIMESTAMP WITH TIME ZONE NOT NULL,
        entity_count  INT NOT NULL,
        added_count   INT NOT NULL,
        removed_count INT NOT NULL,
        changed_count INT NOT NULL
);

CREATE UNIQUE INDEX list_refs_list_hash_idx ON list_refs (list_name, list_hash);

CREATE INDEX list_refs_list_id_idx ON list_refs (list_name, id);

CREATE TABLE list_entity_versions(
        id          BIGSERIAL PRIMARY KEY,
        list_name   VARCHAR(64) NOT NULL,
        source_id   VARCHAR(255) NOT NULL,
        entity_type VARCHAR(20) NOT NULL,
        name        VARCHAR(1024) NOT NULL,
        entity_hash VARCHAR(64) NOT NULL,
        entity      JSONB NOT NULL,
        from_ref_id BIGINT NOT NULL,
        to_ref_id   BIGINT NULL
);

CREATE INDEX list_entity_versions_window_idx ON list_entity_versions (list_name, from_ref_id, to_ref_id);

CREATE INDEX list_entity_versions_live_idx ON list_entity_versions (list_name, to_ref_id, source_id);

CREATE INDEX list_entity_versions_source_idx ON list_entity_versions (list_name, source_id, from_ref_id);

CREATE UNIQUE INDEX list_entity_versions_live_source_uq ON list_entity_versions (list_name, source_id) WHERE to_ref_id IS NULL;
