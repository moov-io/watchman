CREATE TABLE ingested_entities(
       pk_id      INT NOT NULL AUTO_INCREMENT,
       type       VARCHAR(20) NOT NULL,
       source     VARCHAR(30) NOT NULL,
       source_id  VARCHAR(40) NOT NULL,
       entity     JSON NOT NULL,

       CONSTRAINT ingested_entities_pk PRIMARY KEY (pk_id),
       CONSTRAINT ingested_entities_id_uq UNIQUE (source, source_id),
       INDEX ingested_entities_idx_status (source, source_id ASC)
);
