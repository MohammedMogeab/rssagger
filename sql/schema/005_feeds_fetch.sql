-- +goose up

ALTER TABLE feeds ADD COLUMN last_fetch_at TIMESTAMP ;

-- +goose down
drop COLUMN last_fetch_at from feeds;