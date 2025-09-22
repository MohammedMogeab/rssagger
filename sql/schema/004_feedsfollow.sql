-- +goose up

create table feedsfollow(
    id UUID PRIMARY KEY ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    feed_id  UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE(feed_id,user_id)
   
);

-- +goose down
drop table if exists feedsfollow;