-- +goose up

create table feeds(
    id UUID PRIMARY KEY ,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    url  TEXT Unique NOT NULL,
    user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
   
);

-- +goose down
drop table if exists feeds;