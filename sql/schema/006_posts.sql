-- +goose Up
CREATE TABLE posts (
    id            UUID PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    url           TEXT UNIQUE NOT NULL,
    title         TEXT,                 -- nullable (some feeds omit)
    description   TEXT,                 -- nullable (important!)
    published_at  TIMESTAMPTZ,          -- nullable
    feed_id       UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;
