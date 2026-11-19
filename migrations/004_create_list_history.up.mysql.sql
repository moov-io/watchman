CREATE TABLE list_refs(
       id            BIGINT NOT NULL AUTO_INCREMENT,
       list_name     VARCHAR(64) NOT NULL,
       list_hash     VARCHAR(64) NOT NULL,
       downloaded_at DATETIME NOT NULL,
       entity_count  INT NOT NULL,
       added_count   INT NOT NULL,
       removed_count INT NOT NULL,
       changed_count INT NOT NULL,

       CONSTRAINT list_refs_pk PRIMARY KEY (id),
       CONSTRAINT list_refs_list_hash_uq UNIQUE (list_name, list_hash),
       INDEX list_refs_list_id_idx (list_name, id)
);

CREATE TABLE list_entity_versions(
       id          BIGINT NOT NULL AUTO_INCREMENT,
       list_name   VARCHAR(64) NOT NULL,
       source_id   VARCHAR(255) NOT NULL,
       entity_type VARCHAR(20) NOT NULL,
       name        VARCHAR(1024) NOT NULL,
       entity_hash VARCHAR(64) NOT NULL,
       entity      JSON NOT NULL,
       from_ref_id BIGINT NOT NULL,
       to_ref_id   BIGINT NULL,
       live_source_id VARCHAR(255) AS (IF(to_ref_id IS NULL, source_id, NULL)) STORED,

       CONSTRAINT list_entity_versions_pk PRIMARY KEY (id),
       CONSTRAINT list_entity_versions_live_source_uq UNIQUE (list_name, live_source_id),
       INDEX list_entity_versions_window_idx (list_name, from_ref_id, to_ref_id),
       INDEX list_entity_versions_live_idx (list_name, to_ref_id, source_id),
       INDEX list_entity_versions_source_idx (list_name, source_id, from_ref_id)
);
