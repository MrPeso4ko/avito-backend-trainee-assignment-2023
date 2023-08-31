CREATE SCHEMA IF NOT EXISTS segments_manager;

CREATE TABLE IF NOT EXISTS segments_manager.segments
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS segments_manager.belongs_segment
(
    user_id    BIGINT,
    segment_id BIGINT REFERENCES segments_manager.segments (id),
    UNIQUE (user_id, segment_id)
);