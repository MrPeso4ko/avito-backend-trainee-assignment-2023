CREATE SCHEMA IF NOT EXISTS segments_manager;

CREATE TABLE IF NOT EXISTS segments_manager.segments
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS segments_manager.users
(
    id BIGSERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS segments_manager.belongs_segment
(
    user_id    BIGINT REFERENCES segments_manager.users (id),
    segment_id BIGINT REFERENCES segments_manager.segments (id),
    UNIQUE (user_id, segment_id)
);